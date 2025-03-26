package application

import (
	"context"
	"testing"

	"user-service/internal/domain"
	"user-service/internal/infrastructure"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.User{})
	return db
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB()
	repo := infrastructure.NewUserRepository(db)

	user := &domain.User{
		ID:       "123",
		Email:    "test@example.com",
		Password: "hashedpassword",
		Name:     "Test User",
	}

	err := repo.Create(ctx, user)
	assert.Nil(t, err)

	foundUser, err := repo.FindByEmail(ctx, "test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, "Test User", foundUser.Name)
}
