package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Stats ...
	Stats struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		PlayerID  primitive.ObjectID `bson:"playerID,omitempty"`
		Point     int                `bson:"point,omitempty"`
		TotalGame int                `bson:"totalGame,omitempty"`
		WinGame   int                `bson:"winGame,omitempty"`
		WinRate   float32            `bson:"winRate,omitempty"`
		CreatedAt time.Time          `bson:"createdAt,omitempty"`
		UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	}
)
