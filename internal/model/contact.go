package model

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func (c *Contact) Validate() error {
	return validate.Struct(c)
}

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
}
