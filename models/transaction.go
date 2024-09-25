package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

func (p *Status) Scan(value interface{}) error {
	*p = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

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
	PromoCodeID        uint                `json:"promo_code_id" gorm:"default:null"`
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
	ValidFrom          datatypes.Date `json:"valid_from"`
	ValidUntil         datatypes.Date `json:"valid_until"`
	Status             Status         `json:"status" gorm:"type:status"`
}
