package models

import (
	"github.com/go-playground/validator/v10"
)

type ProductTypeRequest struct {
	Name string `json:"name"`
}

type ProductRequest struct {
	ProductTypeID uint   `json:"product_type_ID" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Price         int    `json:"price" validate:"required"`
}

func (pt *ProductTypeRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(pt)

	return err
}

func (p *ProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	return err
}
