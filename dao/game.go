package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
		statsCol = database.StatsCol()
		ctx      = context.Background()
	)
	count, err := statsCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}
	return int(count)
}
