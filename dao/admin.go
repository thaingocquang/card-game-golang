package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Admin struct{}

// Create ...
func (a Admin) Create() error {
	// default admin
	admin := bson.M{"username": "admin", "password": "123456"}

	if _, err := playerCol.InsertOne(context.Background(), admin); err != nil {
		return err
	}

	return nil
}