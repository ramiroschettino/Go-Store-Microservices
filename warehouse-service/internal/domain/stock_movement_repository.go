package domain

type StockMovementRepository interface {
	GetById(id uint) (*StockMovement, error)
	GetAll() ([]StockMovement, error)
	Create(movement *StockMovement) error
	Update(movement *StockMovement) error
	Delete(id uint) error
}
