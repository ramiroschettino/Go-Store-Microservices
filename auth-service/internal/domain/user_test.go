package domain_test

import (
    "testing"
    "auth-service/internal/domain"
    "github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
    t.Run("debería hashear la contraseña correctamente", func(t *testing.T) {
        user := domain.User{
            Password: "mySecurePassword123",
        }

        err := user.HashPassword()
        
        assert.NoError(t, err)
        assert.NotEqual(t, "mySecurePassword123", user.Password)
        assert.True(t, len(user.Password) > 0)
    })

    t.Run("debería fallar con contraseña vacía", func(t *testing.T) {
        user := domain.User{Password: ""}
        err := user.HashPassword()
        assert.Error(t, err)
    })
}

func TestUser_CheckPassword(t *testing.T) {
    t.Run("debería verificar contraseña correcta", func(t *testing.T) {
        originalPass := "correctPassword123"
        user := domain.User{Password: originalPass}
        _ = user.HashPassword()

        err := user.CheckPassword(originalPass)
        assert.NoError(t, err)
    })

    t.Run("debería fallar con contraseña incorrecta", func(t *testing.T) {
        user := domain.User{Password: "originalPass"}
        _ = user.HashPassword()

        err := user.CheckPassword("wrongPassword")
        assert.Error(t, err)
    })
}