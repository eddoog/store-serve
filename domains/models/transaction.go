package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt    `gorm:"index"`
	Items     []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	ID            uint           `gorm:"primaryKey"`
	TransactionID uint           `gorm:"not null"`
	ProductID     uint           `gorm:"not null"`
	Quantity      int            `gorm:"not null"`
	Price         float64        `gorm:"not null"`
	Subtotal      float64        `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
