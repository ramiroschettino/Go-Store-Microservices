package repository

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestJWTRepository(t *testing.T) {
    secret := "clave_secreta"
    repo := NewJWTRepository(secret)

    t.Run("Generar y verificar token", func(t *testing.T) {
        token, err := repo.GenerateAccessToken("user123")
        assert.NoError(t, err)
        
        userID, err := repo.VerifyAccessToken(token)
        assert.NoError(t, err)
        assert.Equal(t, "user123", userID)
    })

    t.Run("Token inválido", func(t *testing.T) {
        _, err := repo.VerifyAccessToken("token.invalido.123")
        assert.Error(t, err)
    })
}