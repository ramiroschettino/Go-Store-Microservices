package infrastructure

import (
	"context"
	"user-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}
