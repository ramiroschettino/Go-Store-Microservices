package domain

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}