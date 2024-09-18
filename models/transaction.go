package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type Transaction struct {
	ID                 uint                `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	DeletedAt          gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	UserID             uint                `json:"user_id"`
	TotalAmount        float64             `json:"total_amount"`
	Status             string              `json:"status"`
	PaymentMethod      string              `json:"payment_method"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	PromoCodeID        uint                `json:"promo_code_id"`
	PromoCode          PromoCode           `json:"promo_code"`
}

type TransactionDetail struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	TransactionID uint           `json:"transaction_id"`
	ProductID     uint           `json:"product_id"`
	Quantity      uint           `json:"quantity"`
	Amount        float64        `json:"amount"`
}

type PromoCode struct {
	ID                 uint           `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DiscountPercentage float64        `json:"discount_percentage"`
	ValidFrom          time.Time      `json:"valid_from"`
	ValidUntil         time.Time      `json:"valid_until"`
	Status             Status         `json:"status" gorm:"type:enum('active', 'inactive')"`
}
