package domain

import "time"

type User struct {
	ID        string    `gorm:"primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // Almacenará el hash bcrypt
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}
