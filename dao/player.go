package dao

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct{}

// FindByID ...
func (p Player) FindByID(ID primitive.ObjectID) (model.Player, error) {
	var player model.Player

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := playerCol.FindOne(context.Background(), filter).Decode(&player); err != nil {
		return model.Player{}, err
	}

	return player, nil
}

// Create ...
func (p Player) Create(player dto.Player) error {
	// InsertOne
	if _, err := playerCol.InsertOne(context.Background(), player); err != nil {
		return err
	}

	return nil
}

// Update ...
func (p Player) Update(ID primitive.ObjectID, player dto.Player) error {
	update := model.Player{Name: player.Name, Email: player.Email, Password: player.Password}

	// UpdateOne
	if _, err := playerCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": update}); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (p Player) Delete(ID primitive.ObjectID) error {
	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne ...
	if _, err := playerCol.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

// GetList ...
func (p Player) GetList(page, limit int) ([]model.Player, error) {
	var players []model.Player

	// options
	opts := new(options.FindOptions)

	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}

	cursor, err := playerCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &players); err != nil {
		return nil, err
	}

	return players, nil
}
