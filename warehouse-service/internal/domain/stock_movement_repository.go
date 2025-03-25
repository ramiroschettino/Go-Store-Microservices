package domain

import "context"

type StockMovementRepository interface {
	GetByID(ctx context.Context, id uint) (*StockMovement, error)
	GetAll(ctx context.Context) ([]StockMovement, error)
	Create(ctx context.Context, movement *StockMovement) error
	Update(ctx context.Context, movement *StockMovement) error
	Delete(ctx context.Context, id uint) error
	GetMovementsByProduct(ctx context.Context, productID uint) ([]StockMovement, error)
	GetCurrentStock(ctx context.Context, productID uint) (int, error)
}
