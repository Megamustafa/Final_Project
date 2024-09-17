package models

import "github.com/go-playground/validator/v10"

type TransactionRequest struct {
	UserID        uint   `json:"user_id" validate:"required"`
	TotalAmount   uint   `json:"total_amount" validate:"required"`
	Status        string `json:"status" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type TransactionDetailRequest struct {
	TransactionID uint `json:"transaction_id" validate:"required"`
	ProductID     uint `json:"product_id" validate:"required"`
	Quantity      uint `json:"quantity" validate:"required"`
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
