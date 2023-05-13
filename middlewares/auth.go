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
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			return echo.ErrUnauthorized
		}
		db := database.GetRedisClient()
		result := db.Get(context.TODO(), claims.Aid)
		if result.Val() == bearerToken {
			c.Request().Header.Set("user_id", claims.UserId)
			return next(c)
		}
		return echo.ErrUnauthorized

	}
}
