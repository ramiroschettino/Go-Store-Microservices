package repository

import (
	"your_project/internal/domain"

	"gorm.io/gorm"
)

type WarehouseRepositoryImpl struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) domain.WarehouseRepository {
	return &WarehouseRepositoryImpl{db: db}
}

func (r *WarehouseRepositoryImpl) GetByID(id uint) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	if err := r.db.First(&warehouse, id).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *WarehouseRepositoryImpl) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	if err := r.db.Find(&warehouses).Error; err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (r *WarehouseRepositoryImpl) Create(warehouse *domain.Warehouse) error {
	return r.db.Create(warehouse).Error
}

func (r *WarehouseRepositoryImpl) Update(warehouse *domain.Warehouse) error {
	return r.db.Save(warehouse).Error
}

func (r *WarehouseRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Warehouse{}, id).Error
}
