package application

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ramiroschettino/Go-Store-Microservices/checkout-service/api"
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
)

type WarehouseService struct {
	repo           domain.WarehouseRepository
	cache          domain.CacheRepository
	search         domain.SearchService
	checkoutClient api.CheckoutServiceClient
}

func NewWarehouseService(repo domain.WarehouseRepository, cache domain.CacheRepository, search domain.SearchService, checkoutClient api.CheckoutServiceClient) *WarehouseService {
	return &WarehouseService{
		repo:           repo,
		cache:          cache,
		search:         search,
		checkoutClient: checkoutClient,
	}
}

func (s *WarehouseService) BlockStock(ctx context.Context, productID string, quantity int) error {
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return fmt.Errorf("error al convertir productID a int: %v", err)
	}

	req := &api.AddToCartRequest{
		ProductId: int32(productIDInt),
		Quantity:  int32(quantity),
	}

	_, err = s.checkoutClient.AddToCart(ctx, req)
	if err != nil {
		return fmt.Errorf("error llamando a AddToCart de checkout-service: %v", err)
	}

	log.Println("Stock bloqueado correctamente en checkout-service.")
	return nil
}
