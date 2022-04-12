package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

// Create ...
func (a Admin) Create() {
	var adminCol = database.AdminCol()

	count, _ := adminCol.CountDocuments(context.Background(), bson.D{})

	if count == 0 {
		admin := model.Admin{
			ID:       primitive.NewObjectID(),
			Username: "admin",
			Password: "123456",
		}
		adminCol.InsertOne(context.Background(), admin)
	}
}

// FindByUsername ...
func (a Admin) FindByUsername(username string) (model.Admin, error) {
	var (
		adminCol = database.AdminCol()
		admin    model.Admin
	)

	// find player
	filter := bson.M{"username": username}
	err := adminCol.FindOne(context.Background(), filter).Decode(&admin)

	// if err
	if err != nil {
		return admin, err
	}

	return admin, nil
}
