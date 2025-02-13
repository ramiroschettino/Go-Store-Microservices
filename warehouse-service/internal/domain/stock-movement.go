package domain

import "time"

type StockMovement struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}
