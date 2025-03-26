package main

import (
	"context"
	"log"
	"net"

	"tu-proyecto/auth-service/internal/auth"
	pb "tu-proyecto/auth-service/internal/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Aquí validarías usuario/contraseña contra tu DB
	token, err := auth.GenerateToken("usuario123") // Cambia esto!
	return &pb.LoginResponse{Token: token}, err
}

func (s *server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	userID, err := auth.ValidateToken(req.Token)
	return &pb.ValidateResponse{
		Valid:  err == nil,
		UserId: userID,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Println("Auth service running on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
