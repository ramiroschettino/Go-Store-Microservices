package main

import (
	"log"
	"warehouse-service/internal/infrastructure/grpc"
)

func main() {
	log.Println("Iniciando el servidor gRPC...")
	grpc.StartGRPCServer()
}
