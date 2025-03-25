package domain

import "context"

type OrderRepository interface {
    Create(ctx context.Context, order *Order) error
    GetByID(ctx context.Context, id uint) (*Order, error)
    GetByCustomer(ctx context.Context, customer string) ([]*Order, error)
    UpdateStatus(ctx context.Context, id uint, status string) error
    Cancel(ctx context.Context, id uint) error
}