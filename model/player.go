package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Player ...
	Player struct {
		ID        primitive.ObjectID `bson:"_id"`
		Name      string             `bson:"name"`
		Email     string             `bson:"email"`
		Password  string             `bson:"password"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}
)
