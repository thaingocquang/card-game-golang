package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Game struct{}

func (g Game) Create(game model.Game) error {
	var gameCol = database.GameCol()

	// InsertOne
	if _, err := gameCol.InsertOne(context.Background(), game); err != nil {
		return err
	}

	return nil
}

// CountAllGame ...
func (g Game) CountAllGame() int {
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
func (g Game) GetList(page, limit int) ([]model.Game, error) {
	var (
		gameCol = database.GameCol()
		games   []model.Game
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

	cursor, err := gameCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return games, err
	}

	if err = cursor.All(context.Background(), &games); err != nil {
		return nil, err
	}

	return games, nil
}
