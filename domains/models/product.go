package models

import "time"

type ProductCategory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}

type Product struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"type:varchar(255);not null" json:"name"`
	Description string          `gorm:"type:text" json:"description"`
	CategoryID  uint            `gorm:"not null" json:"category_id"` // Foreign key
	Category    ProductCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Price       float64         `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int             `gorm:"not null" json:"stock"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Product) TableName() string {
	return "product"
}
