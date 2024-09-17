package models

import "github.com/go-playground/validator/v10"

type AquacultureFarmsRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	FarmID uint `json:"farm_id" validate:"required"`
}

type FarmTypeRequest struct {
	Name string `json:"name" validate:"required"`
}

type FarmRequest struct {
	FarmTypeID  uint   `json:"farm_type_id" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}

func (af *AquacultureFarmsRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(af)

	return err
}

func (ft *FarmTypeRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(ft)

	return err
}

func (f *FarmRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(f)

	return err
}
