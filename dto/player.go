package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	// Player ...
	Player struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// PlayerLogin ...
	PlayerLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// PlayerUpdate ...
	PlayerUpdate struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}
)

// Validate ...
func (p Player) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 256)),
		validation.Field(&p.Email, validation.Required, is.Email, validation.Length(8, 256)),
		validation.Field(&p.Password, validation.Required, validation.Length(6, 256)),
	)
}

// Validate ...
func (p PlayerLogin) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, is.Email, validation.Length(8, 256)),
		validation.Field(&p.Password, validation.Required, validation.Length(6, 256)),
	)
}

// Validate ...
func (p PlayerUpdate) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Length(3, 256)),
		validation.Field(&p.Email, is.Email, validation.Length(8, 256)),
	)
}
