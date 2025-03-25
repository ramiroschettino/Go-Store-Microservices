package repository

import (
	"context"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
)

type StockMovementRepositoryImpl struct {
	db *gorm.DB
}

func NewStockMovementRepository(db *gorm.DB) domain.StockMovementRepository {
	return &StockMovementRepositoryImpl{db: db}
}

func (r *StockMovementRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.StockMovement, error) {
	var movement domain.StockMovement
	err := r.db.WithContext(ctx).First(&movement, id).Error
	return &movement, err
}

func (r *StockMovementRepositoryImpl) GetAll(ctx context.Context) ([]domain.StockMovement, error) {
	var movements []domain.StockMovement
	err := r.db.WithContext(ctx).Find(&movements).Error
	return movements, err
}

func (r *StockMovementRepositoryImpl) Create(ctx context.Context, movement *domain.StockMovement) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Registrar el movimiento
		if err := tx.Create(movement).Error; err != nil {
			return err
		}
		
		// 2. Actualizar el stock del producto
		return tx.Model(&domain.Product{}).
			Where("id = ?", movement.ProductID).
			Update("stock", gorm.Expr("stock + ?", movement.Quantity)).
			Error
	})
}

func (r *StockMovementRepositoryImpl) Update(ctx context.Context, movement *domain.StockMovement) error {
	return r.db.WithContext(ctx).Save(movement).Error
}

func (r *StockMovementRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.StockMovement{}, id).Error
}

func (r *StockMovementRepositoryImpl) GetMovementsByProduct(ctx context.Context, productID uint) ([]domain.StockMovement, error) {
	var movements []domain.StockMovement
	err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Order("created_at desc").
		Find(&movements).Error
	return movements, err
}

func (r *StockMovementRepositoryImpl) GetCurrentStock(ctx context.Context, productID uint) (int, error) {
	var stock struct{ Stock int }
	err := r.db.WithContext(ctx).
		Model(&domain.Product{}).
		Select("stock").
		Where("id = ?", productID).
		Scan(&stock).Error
	return stock.Stock, err
}