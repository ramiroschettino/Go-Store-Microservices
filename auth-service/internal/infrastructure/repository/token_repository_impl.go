package repository

import (
	"time"
	"auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JWTRepository struct {
	secretKey string
}

func NewJWTRepository(secret string) domain.TokenRepository {
	return &JWTRepository{secretKey: secret}
}

func (r *JWTRepository) GenerateAccessToken(userID string) (string, error) {
	claims := domain.TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(r.secretKey))
}

func (r *JWTRepository) VerifyAccessToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.secretKey), nil
	})

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", err
}

func (r *JWTRepository) GenerateRefreshToken() (string, error) {
	// Implementación similar con mayor tiempo de expiración
	return "", nil
}