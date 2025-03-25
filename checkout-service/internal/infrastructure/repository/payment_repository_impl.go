package repository

import (
	"context"
	"time"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/checkout-service/internal/domain"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) domain.PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

func (r *PaymentRepositoryImpl) Create(ctx context.Context, payment *domain.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *PaymentRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.WithContext(ctx).First(&payment, id).Error
	return &payment, err
}

func (r *PaymentRepositoryImpl) GetByOrder(ctx context.Context, orderID uint) ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&payments).Error
	return payments, err
}

func (r *PaymentRepositoryImpl) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&domain.Payment{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}

func (r *PaymentRepositoryImpl) ProcessPayment(ctx context.Context, orderID uint, amount float64, method string) (*domain.Payment, error) {
	payment := &domain.Payment{
		OrderID: orderID,
		Amount:  amount,
		Method:  method,
		Status:  "pending",
		Date:    time.Now(),
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Crear registro de pago
		if err := tx.Create(payment).Error; err != nil {
			return err
		}
		
		// 2. Simular procesamiento de pago (en producción usaría un gateway real)
		payment.Status = "completed"
		
		// 3. Actualizar estado del pago
		if err := tx.Save(payment).Error; err != nil {
			return err
		}
		
		// 4. Actualizar estado del pedido
		return tx.Model(&domain.Order{}).
			Where("id = ?", orderID).
			Update("status", "paid").
			Error
	})

	return payment, err
}