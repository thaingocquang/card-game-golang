package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	// Admin ...
	Admin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// Validate ...
func (a Admin) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Password, validation.Required, validation.Length(6, 256)),
	)
}
