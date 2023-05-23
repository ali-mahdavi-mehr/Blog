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
	"os"
	"strconv"
	"strings"
	"time"
)

func generateAid() string {
	id := uuid.New()
	return id.String()
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
	go redisDB.Set(context.Background(), generatedAid, accessToken, -1)
	return accessToken, refreshToken, err

}

func ExpireToken(token string) (int64, error) {
	claims := models.Auth{}
	bearerToken := strings.Replace(token, "Bearer ", "", 1)
	_, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || claims.TokenType != "refresh_type" {
		return 0, errors.New("to get new token must send  refresh token")
	}
	if err != nil {
		return 0, errors.New("refresh token expired")
	}

	redisDB := database.GetRedisClient()
	result := redisDB.Get(context.TODO(), claims.Aid)
	if result.Val() == "" {
		return 0, errors.New("refresh token expired")
	}
	redisDB.Del(context.TODO(), claims.Aid)
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
