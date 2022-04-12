package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Bot ...
	Bot struct {
		ID           primitive.ObjectID `bson:"_id"`
		Name         string             `bson:"name"`
		TotalPoints  int                `bson:"totalPoints"`
		RemainPoints int                `bson:"remainPoints"`
		MinBet       int                `bson:"minBet"`
		MaxBet       int                `bson:"maxBet"`
		CreatedAt    time.Time          `bson:"createdAt"`
		UpdatedAt    time.Time          `bson:"updatedAt"`
	}
)
