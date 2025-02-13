package domain

type ProductRepository interface {
	GetById(id uint) (*Product, error)
	GetAll() ([]Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(id uint) error
}
