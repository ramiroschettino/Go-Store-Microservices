package repository

import (
	"user-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}
