package models

import "github.com/go-playground/validator/v10"

type ArticleRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func (a *ArticleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(a)

	return err
}
