package models

import "time"

type UserRole string

const (
	UserRoleAdmin    UserRole = "Admin"
	UserRoleCustomer UserRole = "Customer"
)

type User struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Email    string   `json:"email" gorm:"uniqueIndex;not null"`
	Password string   `json:"-" gorm:"not null"`
	Role     UserRole `json:"role" gorm:"type:varchar(20);not null"`
	Phone    string   `json:"phone" gorm:"not null"`
	IsActive bool     `json:"is_active" gorm:"default:true"`
	TimeStamps

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"refresh_tokens,omitempty"`
	Orders        []Order        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"orders,omitempty"`
	Cart          Cart           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"cart,omitzero"`
}

type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Token     string    `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	TimeStamps

	User User `gorm:"foreignKey:UserID" json:"-"`
}
