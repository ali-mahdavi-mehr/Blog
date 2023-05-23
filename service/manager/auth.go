package manager

import (
	"context"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/service/compiles"
	"github.com/alima12/Blog-Go/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type AuthenticationService struct {
	compiles.UnimplementedAuthenticationServer
}

func (auth *AuthenticationService) Login(ctx context.Context, request *compiles.LoginRequest) (*compiles.Token, error) {
	email := request.Email
	password := request.Password
	db := database.GetDB()
	var user models.User
	db.Model(&models.User{}).Find(&user, "email = ?", email)
	if !utils.CheckPassword(password, user.Password) {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}
	userID := strconv.FormatUint(uint64(user.ID), 10)
	accessToken := utils.CreateToken("access", userID)
	refreshToken := utils.CreateToken("refresh", userID)
	return &compiles.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (auth *AuthenticationService) RefreshToken(ctx context.Context, request *compiles.RefreshTokenRequest) (*compiles.Token, error) {
	return nil, nil
}
