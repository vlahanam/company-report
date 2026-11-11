package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateUserReqDTO struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	FullName *string `json:"full_name"`
}

func (u CreateUserReqDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Password, validation.Max(100)),
	)
}