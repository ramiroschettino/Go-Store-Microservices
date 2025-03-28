package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ramiroschettino/go-store-microservice/auth-service/internal/domain" 
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, err
	}

	return db, nil
}