package entities

type SearchProduct struct {
	Category    string  `json:"category"`
	CategoryIDs []uint  `json:"category_ids"`
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	Search      string  `json:"search"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CategoryID  uint             `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Price       float64          `json:"price"`
	Stock       int              `json:"stock"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
}

type SearchProductResponse struct {
	TotalItems int64             `json:"total_items"`
	Data       []ProductResponse `json:"data"`
}
