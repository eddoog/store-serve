package entities

type AddToCartRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

type CartResponse struct {
	Total float64            `json:"total"`
	Items []CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
