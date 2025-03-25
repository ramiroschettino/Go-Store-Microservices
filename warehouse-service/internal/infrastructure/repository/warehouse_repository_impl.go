package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/ramiroschettino/Go-Store-Microservices/warehouse-service/internal/domain"
)

type WarehouseRepositoryImpl struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) domain.WarehouseRepository {
	return &WarehouseRepositoryImpl{db: db}
}

func (r *WarehouseRepositoryImpl) GetByID(ctx context.Context, id uint) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &warehouse, err
}

func (r *WarehouseRepositoryImpl) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	err := r.db.WithContext(ctx).Find(&warehouses).Error
	return warehouses, err
}

func (r *WarehouseRepositoryImpl) Create(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *WarehouseRepositoryImpl) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Save(warehouse).Error
}

func (r *WarehouseRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Warehouse{}, id).Error
}

func (r *WarehouseRepositoryImpl) GetByLocation(ctx context.Context, location string) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.WithContext(ctx).Where("location = ?", location).First(&warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &warehouse, err
}