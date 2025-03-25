package domain

import "context"

type ProductRepository interface {
	GetByID(ctx context.Context, id uint) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uint) error
	// Métodos específicos
	GetBySKU(ctx context.Context, sku string) (*Product, error)
	GetLowStock(ctx context.Context, threshold int) ([]Product, error)
}
