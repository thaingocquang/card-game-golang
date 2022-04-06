package dao

import (
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Admin struct{}

// Create ...
func (a Admin) Create() error {
	var adminCol = database.AdminCol()

	// default admin
	admin := bson.M{"username": "admin", "password": "123456"}

	if _, err := adminCol.InsertOne(context.Background(), admin); err != nil {
		return err
	}

	return nil
}