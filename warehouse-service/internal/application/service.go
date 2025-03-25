package application

import (
	"context"
	
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
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
	// Implementar lógica de negocio
}