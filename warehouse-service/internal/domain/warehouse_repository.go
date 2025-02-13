package domain

type WarehouseRepository interface {
	GetByID(id uint) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	Create(warehouse *Warehouse) error
	Update(warehouse *Warehouse) error
	Delete(id uint) error
}
