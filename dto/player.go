package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Name     string `json:"name,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	// Profile ...
	Profile struct {
		ID        primitive.ObjectID `json:"id"`
		Name      string             `json:"name"`
		Email     string             `json:"email"`
		Point     int                `json:"point"`
		TotalGame int                `json:"totalGame"`
		WinGame   int                `json:"winGame"`
		WinRate   float32            `json:"winRate"`
	}

	ProfileUpdate struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Point    int    `json:"point"`
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

// Validate ...
func (pu ProfileUpdate) Validate() error {
	return validation.ValidateStruct(&pu,
		validation.Field(&pu.Name, validation.Length(3, 256)),
		validation.Field(&pu.Email, is.Email, validation.Length(8, 256)),
		validation.Field(&pu.Password, validation.Length(8, 256)),
		validation.Field(&pu.Point,
			validation.Min(0),
			validation.Max(10000)),
	)
}
