package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=user-db user=admin password=secret dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a PostgreSQL: %v", err)
	}
	log.Println("Conectado a PostgreSQL")
	return db
}
