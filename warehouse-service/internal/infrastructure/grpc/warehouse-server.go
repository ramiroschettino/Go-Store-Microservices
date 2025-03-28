package grpc

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	pb "github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/api"
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/application"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedWarehouseServiceServer
	warehouseService application.WarehouseService
}

func NewServer(warehouseService application.WarehouseService) *Server {
	return &Server{
		warehouseService: warehouseService,
	}
}

func (s *Server) CheckAndBlockStock(ctx context.Context, req *pb.StockRequest) (*pb.StockResponse, error) {
	// Validación de entrada
	if req.GetProductId() == "" || req.GetQuantity() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "product_id y quantity son requeridos")
	}

	// Timeout context
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Llamada al servicio de aplicación
	available, err := s.warehouseService.CheckAndBlockStock(ctx, req.GetProductId(), int(req.GetQuantity()))
	if err != nil {
		if errors.Is(err, application.ErrInsufficientStock) {
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "error al verificar stock: %v", err)
	}

	return &pb.StockResponse{
		Success:           true,
		AvailableQuantity: int32(available),
		Message:           "Stock reservado exitosamente",
	}, nil
}

func (s *Server) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.StockResponse, error) {
	// Validación de entrada
	if req.GetProductId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "product_id es requerido")
	}

	// Timeout context
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Llamada al servicio de aplicación
	available, err := s.warehouseService.UpdateStock(
		ctx,
		req.GetProductId(),
		int(req.GetQuantityToDeduct()),
		req.GetRevert(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error al actualizar stock: %v", err)
	}

	return &pb.StockResponse{
		Success:           true,
		AvailableQuantity: int32(available),
		Message:           "Stock actualizado exitosamente",
	}, nil
}

func StartGRPCServer(port string, warehouseService application.WarehouseService) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	pb.RegisterWarehouseServiceServer(s, NewServer(warehouseService))

	log.Printf("Servidor gRPC iniciado en %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()

	log.Printf("Iniciando llamada: %s", info.FullMethod)

	resp, err = handler(ctx, req)

	log.Printf("Llamada completada: %s, Duración: %v, Error: %v",
		info.FullMethod,
		time.Since(start),
		err)

	return resp, err
}
