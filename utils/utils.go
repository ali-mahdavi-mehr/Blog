package utils

import (
	"context"
	"fmt"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

func generateAid() string {
	id := uuid.New()
	return id.String()
}

func CreateToken(tokenType, username string) string {
	var JwtExpireTime int64
	switch tokenType {
	case "access":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("jwt_access_expire_time"), 10, 32)
	case "refresh":
		JwtExpireTime, _ = strconv.ParseInt(os.Getenv("jwt_refresh_expire_time"), 10, 32)

	}
	generatedAid := generateAid()
	claims := &models.Auth{
		username,
		generatedAid,
		fmt.Sprintf("%s_type", tokenType),

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(JwtExpireTime))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	createdToken, err := token.SignedString([]byte(os.Getenv("secret_key")))
	if err != nil {
		return ""
	}
	redisDB := database.GetRedisClient()
	go redisDB.Set(context.Background(), generatedAid, createdToken, -1)
	return createdToken

}
