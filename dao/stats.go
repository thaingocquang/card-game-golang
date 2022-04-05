package dao

import (
	"card-game-golang/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Stats struct{}

// FindByID ...
func (p Stats) FindByID(ID primitive.ObjectID) (model.Stats, error) {
	var stats model.Stats

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := statsCol.FindOne(context.Background(), filter).Decode(&stats); err != nil {
		return model.Stats{}, err
	}

	return stats, nil
}

// Create ...
func (p Stats) Create(playerID primitive.ObjectID) error {
	// default stats when create player
	var stats = model.Stats{
		ID:        primitive.NewObjectID(),
		PlayerID:  playerID,
		Point:     0,
		TotalGame: 0,
		WinGame:   0,
		WinRate:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// InsertOne
	if _, err := statsCol.InsertOne(context.Background(), stats); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (p Stats) Delete(playerID primitive.ObjectID) error {
	// filter delete by playerID
	filter := bson.M{"playerID": playerID}

	// DeleteOne ...
	if _, err := statsCol.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}
