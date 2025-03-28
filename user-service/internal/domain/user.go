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
