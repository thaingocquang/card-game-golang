package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// BotJSON ...
	BotJSON struct {
		ID           primitive.ObjectID `json:"id"`
		Name         string             `json:"name"`
		TotalPoints  int                `json:"totalPoints"`
		RemainPoints int                `json:"remainPoints"`
		MinBet       int                `json:"minBet"`
		MaxBet       int                `json:"maxBet"`
	}

	// Bot ...
	Bot struct {
		Name         string `json:"name"`
		TotalPoints  int    `json:"totalPoints"`
		RemainPoints int    `json:"remainPoints"`
		MinBet       int    `json:"minBet"`
		MaxBet       int    `json:"maxBet"`
	}
)

// ValidateCreate ...
func (b Bot) ValidateCreate() error {
	return validation.ValidateStruct(&b,
		// validate field Name
		validation.Field(&b.Name,
			validation.Required,
			validation.Length(1, 256)),
		// validate field TotalPoints
		validation.Field(&b.TotalPoints,
			validation.Required,
			validation.Min(0),
			validation.Max(10000)),
		// validate field MinBet
		validation.Field(&b.MinBet,
			validation.Required,
			validation.Min(0),
			validation.Max(10000)),
		// validate field MaxBet
		validation.Field(&b.MaxBet,
			validation.Required,
			validation.Min(0),
			validation.Max(10000)),
	)
}

// ValidateUpdate ...
func (b Bot) ValidateUpdate() error {
	return validation.ValidateStruct(&b,
		// validate field Name
		validation.Field(&b.Name,
			validation.Length(1, 256)),
		// validate field TotalPoints
		validation.Field(&b.TotalPoints,
			validation.Min(0),
			validation.Max(10000)),
		validation.Field(&b.RemainPoints,
			validation.Min(0),
			validation.Max(10000)),
		// validate field MinBet
		validation.Field(&b.MinBet,
			validation.Min(0),
			validation.Max(10000)),
		// validate field MaxBet
		validation.Field(&b.MaxBet,
			validation.Min(0),
			validation.Max(10000)),
	)
}
