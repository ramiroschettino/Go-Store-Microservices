package repository

import (
	"context"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/checkout-service/internal/domain"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *OrderRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *OrderRepositoryImpl) GetByCustomer(ctx context.Context, customer string) ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.WithContext(ctx).Where("customer = ?", customer).Find(&orders).Error
	return orders, err
}

func (r *OrderRepositoryImpl) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&domain.Order{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}

func (r *OrderRepositoryImpl) Cancel(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Order{}).
			Where("id = ?", id).
			Update("status", "cancelled").Error; err != nil {
			return err
		}
		
		// Aquí podrías agregar lógica para revertir el stock
		return nil
	})
}