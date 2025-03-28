package application

import (
	"context"
	
	"warehouse-service/internal/domain"
)

type WarehouseService struct {
	repo      domain.WarehouseRepository
	cache     domain.Cache
	search    domain.SearchService
}

func NewWarehouseService(repo domain.WarehouseRepository, cache domain.Cache, search domain.SearchService) *WarehouseService {
	return &WarehouseService{
		repo:      repo,
		cache:     cache,
		search:    search,
	}
}

func (s *WarehouseService) BlockStock(ctx context.Context, productID string, quantity int) error {
	// Aquí realizas la llamada gRPC al servicio checkout
	req := &api.BlockStockRequest{
		ProductId: productID,
		Quantity:  int32(quantity),
	}

	// Llamada al servicio checkout
	_, err := s.checkoutClient.BlockStock(ctx, req)
	if err != nil {
		return fmt.Errorf("error llamando a BlockStock de checkout-service: %v", err)
	}

	// Procesa el bloqueo de stock en tu servicio después de la llamada
	// Código de negocio relacionado con el bloqueo de stock en warehouse-service
	log.Println("Stock bloqueado correctamente en checkout-service.")
	return nil
}