package domain

import "context"

type WarehouseRepository interface {
	GetByID(ctx context.Context, id uint) (*Warehouse, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	Create(ctx context.Context, warehouse *Warehouse) error
	Update(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id uint) error
	// Métodos específicos
	GetByLocation(ctx context.Context, location string) (*Warehouse, error)
}
