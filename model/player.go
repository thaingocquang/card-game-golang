package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Player ...
	Player struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Name      string             `bson:"name,omitempty"`
		Email     string             `bson:"email,omitempty"`
		Password  string             `bson:"password,omitempty"`
		CreatedAt time.Time          `bson:"createdAt,omitempty"`
		UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	}

	Profile struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Stat      Stats              `bson:"stat"`
		Name      string             `bson:"name,omitempty"`
		Email     string             `bson:"email,omitempty"`
		Password  string             `bson:"password,omitempty"`
		CreatedAt time.Time          `bson:"createdAt,omitempty"`
		UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
	}
)
