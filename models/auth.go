package models

import "github.com/golang-jwt/jwt/v4"

type Auth struct {
	UserId    string `json:"user_id"`
	Aid       string `json:"aid"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
