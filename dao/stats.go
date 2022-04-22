package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Stats struct{}

// FindByID ...
func (s Stats) FindByID(ID primitive.ObjectID) (model.Stats, error) {
	var (
		statsCol = database.StatsCol()
		stats    model.Stats
	)

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := statsCol.FindOne(context.Background(), filter).Decode(&stats); err != nil {
		return model.Stats{}, err
	}

	return stats, nil
}

// FindByPlayerID ...
func (s Stats) FindByPlayerID(ID primitive.ObjectID) (model.Stats, error) {
	var (
		statsCol = database.StatsCol()
		stats    model.Stats
	)

	// filter
	filter := bson.M{"playerID": ID}

	// FindOne
	if err := statsCol.FindOne(context.Background(), filter).Decode(&stats); err != nil {
		return model.Stats{}, err
	}

	return stats, nil
}

// Create ...
func (s Stats) Create(playerID primitive.ObjectID) error {
	var statsCol = database.StatsCol()

	// default stats when create player
	var stats = model.Stats{
		ID:        primitive.NewObjectID(),
		PlayerID:  playerID,
		Point:     2000,
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

// DeleteByID ...
func (s Stats) DeleteByID(playerID primitive.ObjectID) error {
	var statsCol = database.StatsCol()

	// filter delete by playerID
	filter := bson.M{"playerID": playerID}

	// DeleteOne ...
	if _, err := statsCol.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

// DeleteAll ...
func (s Stats) DeleteAll() error {
	var statsCol = database.StatsCol()

	// DeleteMany
	if _, err := statsCol.DeleteMany(context.Background(), bson.M{}); err != nil {
		return err
	}

	return nil
}

// GetList ...
func (s Stats) GetList(page, limit int) ([]model.Stats, error) {
	var (
		statsCol = database.StatsCol()
		stats    = make([]model.Stats, 0)
	)

	// options
	opts := new(options.FindOptions)

	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}

	cursor, err := statsCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return stats, err
	}

	if err = cursor.All(context.Background(), &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// UpdateByID ...
func (s Stats) UpdateByID(id primitive.ObjectID, stats model.Stats) error {
	var statsCol = database.StatsCol()

	// UpdateOne
	if _, err := statsCol.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": stats}); err != nil {
		return err
	}

	return nil
}

// UpdateByPlayerID ...
func (s Stats) UpdateByPlayerID(playerID primitive.ObjectID, stats model.StatsUpdate) error {
	var statsCol = database.StatsCol()

	//// UpdateOne
	//if _, err := statsCol.UpdateOne(context.Background(), bson.M{"playerID": playerID}, bson.M{"$set": stats}); err != nil {
	//	return err
	//}

	// UpdateOne
	ur, err := statsCol.UpdateOne(context.Background(), bson.M{"playerID": playerID}, bson.M{"$set": stats})
	if err != nil {
		return err
	}

	fmt.Println(ur.UpsertedCount)

	return nil
}
