package main

import (
	"log"

	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/infrastructure/grpc"
)

func main() {
	log.Println("Iniciando el servidor gRPC...")
	grpc.StartGRPCServer()
}
