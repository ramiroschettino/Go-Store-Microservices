package main

import (
	"log"
	"os"

	"github.com/ramiroschettino/Go-Store-Microservices/auth-service/internal/application"
	"github.com/ramiroschettino/Go-Store-Microservices/auth-service/internal/infrastructure/db"
	"github.com/ramiroschettino/Go-Store-Microservices/auth-service/internal/infrastructure/repository"
)

func main() {
	// Configuración
	dsn := os.Getenv("DB_DSN")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Conexión a la base de datos
	dbConn, err := db.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Repositorios
	userRepo := repository.NewUserRepository(dbConn)
	tokenRepo := repository.NewJWTRepository(jwtSecret)

	// Crear el servicio de autenticación
	authService := application.NewAuthService(userRepo, tokenRepo)

	// Aquí es donde llamas a las funciones de authService para manejar la lógica de negocio
	// Ejemplo de registro de usuario
	email := "test@example.com"
	password := "password123"
	user, err := authService.Register(email, password)
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}
	log.Printf("User registered: %v", user)

	// Ejemplo de login de usuario
	token, err := authService.Login(email, password)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Access token: %v", token.AccessToken)
}
