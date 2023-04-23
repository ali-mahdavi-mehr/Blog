package api

import (
	"context"
	"fmt"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
	"time"
)

func createToken(tokenType, username string) string {
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(JwtExpireTime))),
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

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}
	accessToken := createToken("access", username)
	refreshToken := createToken("refresh", username)
	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func generateAid() string {
	id := uuid.New()
	return id.String()
}
