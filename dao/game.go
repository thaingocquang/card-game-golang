package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"card-game-golang/util"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct{}

func (Game) Create(game model.Game) error {
	var gameCol = database.GameCol()

	// InsertOne
	if _, err := gameCol.InsertOne(context.Background(), game); err != nil {
		return err
	}

	return nil
}

// CountAllGame ...
func (Game) CountAllGame() int {
	var (
		gameCol = database.GameCol()
		ctx     = context.Background()
	)
	count, err := gameCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}
	return int(count)
}

// GetList ...
func (Game) GetList(paging *util.Paging) ([]model.Game, error) {
	var (
		gameCol = database.GameCol()
		games   []model.Game
	)

	// options
	opts := new(options.FindOptions)
	opts.SetSkip(int64((paging.Page - 1) * paging.Limit))
	opts.SetLimit(int64(paging.Limit))

	// count document in playerCol
	count, err := gameCol.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	// set paging total
	paging.Total = count

	cursor, err := gameCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return games, err
	}

	if err = cursor.All(context.Background(), &games); err != nil {
		return nil, err
	}

	return games, nil
}

// RecentByPlayerID ...
func (Game) RecentByPlayerID(ID primitive.ObjectID) ([]model.Game, error) {
	var (
		gameCol = database.GameCol()
		games   []model.Game
	)

	// options
	opts := new(options.FindOptions)

	opts.SetSort(bson.D{{"createdAt", -1}})
	opts.SetLimit(5)

	cursor, err := gameCol.Find(context.Background(), bson.M{"playerID": ID}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &games); err != nil {
		return nil, err
	}

	return games, nil
}
