package dto

import validation "github.com/go-ozzo/ozzo-validation"

type (
	// Player ...
	Player struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (p Player) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Email, validation.Required),
		validation.Field(&p.Password, validation.Required),
	)
}
