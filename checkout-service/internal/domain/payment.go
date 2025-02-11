package domain

import "time"

type Payment struct {
	ID      uint      `json:"id" gorm:"primaryKey"`
	OrderID uint      `json:"order_id" gorm:"not null"`
	Amount  float64   `json:"amount" gorm:"not null"`
	Method  string    `json:"method" gorm:"not null"`
	Status  string    `json:"status" gorm:"not null"`
	Date    time.Time `json:"date" gorm:"not null"`
}
