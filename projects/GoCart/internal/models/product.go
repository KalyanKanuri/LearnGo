package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"uniqueIndex;not null"`
	Description string `json:"description" gorm:"not null"`
	TimeStamps

	// Relations
	Products []Product `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"products,omitempty"`
}

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id" gorm:"index;not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       uint    `json:"stock" gorm:"default:0"`
	SKU         string  `json:"sku" gorm:"uniqueIndex;not null"`
	IsActive    bool    `json:"is_active" gorm:"default:true"`
	TimeStamps

	// Relations
	Category   Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Images     []ProductImage `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"images,omitempty"`
	OrderItems []OrderItem    `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"order_items,omitempty"`
	CartItems  []CartItem     `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"cart_items,omitempty"`
}

type ProductImage struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id" gorm:"not null;index"`
	URL       string `json:"url" gorm:"not null"`
	TimeStamps
}
