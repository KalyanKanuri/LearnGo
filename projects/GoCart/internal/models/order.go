package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusConfirmed OrderStatus = "Confirmed"
	OrderStatusShipped   OrderStatus = "Shipped"
	OrderStatusDelivered OrderStatus = "Delivered"
	OrderStatusCompleted OrderStatus = "Completed"
	OrderStatusCancelled OrderStatus = "Cancelled"
)

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"user_id" gorm:"not null;index"`
	Status     OrderStatus `json:"order_status" gorm:"type:varchar(20);not null"`
	TotalPrice float64     `json:"total_price" gorm:"type:decimal(10,2);not null"`
	TimeStamps

	// Relations
	User       User        `gorm:"foreignKey:UserID" json:"user,omitzero"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null;index"`
	ProductID uint    `json:"product_id" gorm:"not null;index"`
	Quantity  uint    `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TimeStamps

	// Relations
	Order   Order   `gorm:"foreignKey:OrderID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitzero"`
}

type Cart struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	TimeStamps

	// Relations
	CartItems []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE" json:"cart_items,omitempty"`
}

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartID    uint    `json:"cart_id" gorm:"not null;index"`
	ProductID uint    `json:"product_id" gorm:"not null;index"`
	Quantity  uint    `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TimeStamps

	// Relations
	Cart    Cart    `gorm:"foreignKey:CartID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitzero"`
}
