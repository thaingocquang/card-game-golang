package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	// Admin ...
	Admin struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
	}
)
