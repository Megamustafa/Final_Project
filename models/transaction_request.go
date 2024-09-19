package models

import (
	"github.com/go-playground/validator/v10"
)

type TransactionRequest struct {
	UserID        uint   `json:"user_id" validate:"required"`
	TotalAmount   uint   `json:"total_amount" validate:"required"`
	Status        string `json:"status" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	PromoCodeID   uint   `json:"promo_code_id"`
}

type TransactionStatusRequest struct {
	Status        string `json:"status" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type TransactionDetailRequest struct {
	TransactionID uint `json:"transaction_id" validate:"required"`
	ProductID     uint `json:"product_id" validate:"required"`
	Quantity      uint `json:"quantity" validate:"required"`
}

type PromoCodeRequest struct {
	DiscountPercentage float64 `json:"discount_percentage" validate:"required"`
	ValidFrom          string  `json:"valid_from" validate:"required"`
	ValidUntil         string  `json:"valid_until" validate:"required"`
	Status             Status  `json:"status" validate:"required"`
}

func (t *TransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	return err
}

func (td *TransactionDetailRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(td)

	return err
}

func (pc *PromoCodeRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(pc)

	return err
}

func (ts *TransactionStatusRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(ts)

	return err
}
