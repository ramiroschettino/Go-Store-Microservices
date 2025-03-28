package repository

import (
	"auth-service/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

// Crear un nuevo usuario sin contexto
func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// Buscar un usuario por su ID sin contexto
func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

// Buscar un usuario por su correo electrónico sin contexto
func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
