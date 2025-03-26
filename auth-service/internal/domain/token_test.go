package domain_test

import (
    "testing"
    "time"
    "auth-service/internal/domain"
    "github.com/stretchr/testify/assert"
)

func TestTokenClaims_Valid(t *testing.T) {
    t.Run("claims válidas", func(t *testing.T) {
        claims := domain.TokenClaims{
            UserID: "123",
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
            },
        }
        
        err := claims.Valid()
        assert.NoError(t, err)
    })

    t.Run("claims expiradas", func(t *testing.T) {
        claims := domain.TokenClaims{
            UserID: "123",
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
            },
        }
        
        err := claims.Valid()
        assert.Error(t, err)
    })
}