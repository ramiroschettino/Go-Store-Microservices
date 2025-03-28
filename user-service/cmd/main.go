package main

import (
	"log"
	"net"
	"user-service/api"
	"user-service/internal/application"
	"user-service/internal/domain"
	"user-service/internal/infrastructure/db"
	infrastructure "user-service/internal/infrastructure/repository"

	"google.golang.org/grpc"
)

func main() {
	// Database
	dbConn := db.Connect()
	dbConn.AutoMigrate(&domain.User{})

	// Dependencias
	userRepo := infrastructure.NewUserRepository(dbConn)
	userApp := application.NewUserService(userRepo)

	// gRPC Server
	grpcServer := grpc.NewServer()
	api.RegisterUserServiceServer(grpcServer, userApp)

	// Escuchar
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Servidor UserService iniciado en :50054")
	grpcServer.Serve(lis)
}
