package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Stats ...
	Stats struct {
		ID        primitive.ObjectID `bson:"_id"`
		PlayerID  primitive.ObjectID `bson:"playerID"`
		Point     int                `bson:"point"`
		TotalGame int                `bson:"totalGame"`
		WinGame   int                `bson:"winGame"`
		WinRate   float32            `bson:"winRate"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}
)
