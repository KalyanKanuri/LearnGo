package dto

import "time"

type OrderItemResponse struct {
	ID        uint            `json:"id"`
	Product   ProductResponse `json:"product"`
	Quantity  uint            `json:"quantity"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type OrderResponse struct {
	ID        uint                `json:"id"`
	UserID    uint                `json:"user_id"`
	Status    string              `json:"status"`
	Items     []OrderItemResponse `json:"items"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
