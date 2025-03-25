package domain

import "context"

type PaymentRepository interface {
    Create(ctx context.Context, payment *Payment) error
    GetByID(ctx context.Context, id uint) (*Payment, error)
    GetByOrder(ctx context.Context, orderID uint) ([]*Payment, error)
    UpdateStatus(ctx context.Context, id uint, status string) error
    ProcessPayment(ctx context.Context, orderID uint, amount float64, method string) (*Payment, error)
}