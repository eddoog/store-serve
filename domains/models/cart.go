package models

type Cart struct {
	ID     uint       `gorm:"primaryKey" json:"id"`
	UserID uint       `json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Snapshot of product price at the time of addition
}

func (Cart) TableName() string {
	return "cart"
}

func (CartItem) TableName() string {
	return "cart_item"
}
