package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateAid() string {
	id := uuid.New()
	return id.String()
}

func createAuthRedisKey(aid, userId string) string {
	return fmt.Sprintf("%s_%s", aid, userId)
}

func createToken(tokenType, userId, aid string) (string, error) {
	var JwtExpireTime int64
	switch tokenType {
	case "access":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("JWT_ACCESS_EXPIRE_TIME"), 10, 32)
	case "refresh":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("JWT_REFRESH_EXPIRE_TIME"), 10, 32)

	}

	claims := &models.Auth{
		userId,
		aid,
		fmt.Sprintf("%s_type", tokenType),

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(JwtExpireTime))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	createdToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return createdToken, err
}

func CreateTokens(userId string) (string, string, error) {
	generatedAid := generateAid()
	var err error
	var accessToken, refreshToken string
	accessToken, err = createToken("access", userId, generatedAid)
	refreshToken, err = createToken("refresh", userId, generatedAid)
	redisDB := database.GetRedisClient()
	go redisDB.Set(context.Background(), createAuthRedisKey(generatedAid, userId), accessToken, -1)
	return accessToken, refreshToken, err

}

func GetTokenClaims(token string) (*models.Auth, error) {
	claims := models.Auth{}
	bearerToken := strings.Replace(token, "Bearer ", "", 1)
	_, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.New("token expired")
	}
	return &claims, nil

}

func ExpireToken(token string, isRefreshToken bool) (int64, error) {
	claims, err := GetTokenClaims(token)
	if err != nil {
		return 0, err
	}

	if isRefreshToken && claims.TokenType != "refresh_type" {
		return 0, errors.New("to get new token must send  refresh token")
	}
	redisDB := database.GetRedisClient()
	redisKey := createAuthRedisKey(claims.Aid, claims.UserId)
	result := redisDB.Get(context.TODO(), redisKey)
	if result.Val() == "" {
		return 0, errors.New("token already expired")
	}
	redisDB.Del(context.TODO(), redisKey)
	userID, _ := strconv.ParseInt(claims.UserId, 10, 64)
	return userID, nil

}

func ConvertToTimestamp(t time.Time) (*timestamp.Timestamp, error) {
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func IsValidToken(token string) (bool, string, error) {
	claims, err := GetTokenClaims(token)
	bearerToken := strings.Replace(token, "Bearer ", "", 1)
	if err != nil {
		return false, "", errors.New("error in decode token")
	}
	db := database.GetRedisClient()
	result := db.Get(context.TODO(), createAuthRedisKey(claims.Aid, claims.UserId))
	if result.Val() == bearerToken {
		return true, claims.UserId, nil
	}
	return false, "", errors.New("invalid token")
}

func CheckAuthorizationInGRPC(ctx context.Context) error {
	data := metadata.ValueFromIncomingContext(ctx, "access_token")
	if data == nil {
		errMessage := "Unauthorized"
		return status.Error(codes.Unauthenticated, errMessage)
	}
	ok, _, err := IsValidToken(data[0])
	if !ok {
		errMessage := err.Error()
		return status.Error(codes.Unauthenticated, errMessage)
	}
	return nil

}

func RevokeAllTokens(userId string) {
	db := database.GetRedisClient()
	keys, _ := db.Keys(context.TODO(), createAuthRedisKey("*", userId)).Result()
	for _, key := range keys {
		if strings.Contains(key, userId) {
			db.Del(context.TODO(), key)
		}
	}
}
