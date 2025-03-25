package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "warehouse-service/api" // Importa los archivos generados

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWarehouseServiceServer
}

// Implementa el método CheckAndBlockStock
func (s *server) CheckAndBlockStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	// Aquí va la lógica de tu servicio, como bloquear el stock
	fmt.Printf("Verificando stock para el producto: %s, cantidad: %d\n", req.ProductId, req.Quantity)

	// Simulación de respuesta
	return &pb.StockResponse{
		Success:           true,
		AvailableQuantity: 100, // Esto lo puedes cambiar según la lógica
		Message:           "Stock verificado y bloqueado correctamente.",
	}, nil
}

// Implementa el método UpdateStock
func (s *server) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.StockResponse, error) {
	// Lógica para actualizar el stock
	fmt.Printf("Actualizando stock para el producto: %s, cantidad a restar: %d\n", req.ProductId, req.QuantityToDeduct)

	// Simulación de respuesta
	return &pb.StockResponse{
		Success:           true,
		AvailableQuantity: 50, // Esto lo puedes cambiar según la lógica
		Message:           "Stock actualizado correctamente.",
	}, nil
}

// Función para iniciar el servidor gRPC
func StartGRPCServer() {
	// Establece el puerto donde se escuchará el servidor gRPC
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Crea un servidor gRPC
	s := grpc.NewServer()

	// Registra el servidor con los métodos definidos
	pb.RegisterWarehouseServiceServer(s, &server{})

	// Inicia el servidor gRPC
	log.Printf("Servidor gRPC corriendo en %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
