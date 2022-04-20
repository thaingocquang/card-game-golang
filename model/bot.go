package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Bot ...
	Bot struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Name         string             `bson:"name,omitempty"`
		TotalPoints  int                `bson:"totalPoints,omitempty"`
		RemainPoints int                `bson:"remainPoints,omitempty"`
		MinBet       int                `bson:"minBet,omitempty"`
		MaxBet       int                `bson:"maxBet,omitempty"`
		CreatedAt    time.Time          `bson:"createdAt,omitempty"`
		UpdatedAt    time.Time          `bson:"updatedAt,omitempty"`
	}
)
