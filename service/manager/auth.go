package manager

import (
	"context"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/service/compiles"
	"github.com/alima12/Blog-Go/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	accessToken, refreshToken, err := utils.CreateTokens(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &compiles.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (auth *AuthenticationService) RefreshToken(ctx context.Context, request *compiles.RefreshTokenRequest) (*compiles.Token, error) {
	userId, err := utils.ExpireToken(request.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	userID := strconv.FormatUint(uint64(userId), 10)
	accessToken, refreshToken, tokenErr := utils.CreateTokens(userID)
	if tokenErr != nil {
		return nil, status.Error(codes.Internal, tokenErr.Error())
	}
	return &compiles.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (auth *AuthenticationService) Logout(ctx context.Context, request *compiles.Empty) (*compiles.Empty, error) {
	if err := utils.CheckAuthorizationInGRPC(ctx); err != nil {
		return nil, err
	}
	return &compiles.Empty{}, nil
}

func (auth *AuthenticationService) ChangePassword(ctx context.Context, request *compiles.ChangePasswordRequest) (*compiles.Empty, error) {
	if err := utils.CheckAuthorizationInGRPC(ctx); err != nil {
		return nil, err
	}
	if request.Password != request.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "Passwords don't match")
	}
	data := metadata.ValueFromIncomingContext(ctx, "access_token")
	claims, err := utils.GetTokenClaims(data[0])
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if request.ConfirmPassword != request.Password {
		return nil, status.Error(codes.InvalidArgument, "Passwords don't match")
	}

	if request.OldPassword == request.Password {
		return nil, status.Error(codes.InvalidArgument, "Old and new password are the same")
	}

	db := database.GetDB()
	var user models.User
	db.Model(&models.User{}).Find(&user, claims.UserId)
	if !utils.CheckPassword(request.OldPassword, user.Password) {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}
	user.Password, _ = utils.HashPassword(request.Password)
	db.Save(&user)
	return &compiles.Empty{}, nil
}
