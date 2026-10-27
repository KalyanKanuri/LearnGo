package dto

import "time"

type AddToCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity"`
}

type CartItemResponse struct {
	ID        uint            `json:"id"`
	Product   ProductResponse `json:"product"`
	Quantity  uint            `json:"quantity"`
	Subtotal  float64         `json:"sub_total"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type CartResponse struct {
	ID        uint               `json:"id"`
	UserID    uint               `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	Total     float64            `json:"total"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
