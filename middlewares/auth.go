package middlewares

import (
	"context"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
	"os"
	"strings"
)

func LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := models.Auth{}
		bearerToken := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", 1)
		_, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("secret_key")), nil
		})
		if err != nil {
			return echo.ErrBadRequest
		}
		db := database.GetRedisClient()
		result := db.Get(context.TODO(), claims.Aid)
		if result.Val() == bearerToken {
			return next(c)
		}
		return echo.ErrUnauthorized

	}
}
