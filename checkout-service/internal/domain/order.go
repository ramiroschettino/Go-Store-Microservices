package domain

import "time"

type Order struct {
	ID        uint      `json:"id" gorm: "primaryKey"`
	Customer  string    `json:"customer" gorm: "not null"`
	ProductID uint      `json:"product_id" gorm: "not null"`
	Quantity  uint      `json:"quantity" gorm: "not null`
	Total     float64   `json:"total" gorm: "not null"`
	Date      time.Time `json: "date" gorm: "not null"`
}
