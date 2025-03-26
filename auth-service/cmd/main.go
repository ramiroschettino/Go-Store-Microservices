package main

import (
	"auth-service/internal/application"
	"auth-service/internal/infrastructure/db"
	"auth-service/internal/infrastructure/repository"
	"log"
	"os"
)

func main() {
	// Configuración
	dsn := os.Getenv("DB_DSN")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Database
	dbConn, err := db.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Repositorios
	userRepo := repository.NewUserRepository(dbConn)
	tokenRepo := repository.NewJWTRepository(jwtSecret)

	// Servicio
	authService := application.NewAuthService(userRepo, tokenRepo)

	// Iniciar gRPC/HTTP server (implementar después)
	log.Println("Auth service started successfully")
}