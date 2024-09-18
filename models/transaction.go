package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                 uint                `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	DeletedAt          gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	UserID             uint                `json:"user_id"`
	TotalAmount        uint                `json:"total_amount"`
	Status             string              `json:"status"`
	PaymentMethod      string              `json:"payment_method"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	TransactionID uint           `json:"transaction_id"`
	ProductID     uint           `json:"product_id"`
	Quantity      uint           `json:"quantity"`
	Amount        uint           `json:"amount"`
}
