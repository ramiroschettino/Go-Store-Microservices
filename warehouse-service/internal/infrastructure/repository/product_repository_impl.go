package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &product, err
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Product{}, id).Error
}

func (r *ProductRepositoryImpl) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).Where("sku = ?", sku).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &product, err
}

func (r *ProductRepositoryImpl) GetLowStock(ctx context.Context, threshold int) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(ctx).
		Where("stock < ?", threshold).
		Find(&products).Error
	return products, err
}