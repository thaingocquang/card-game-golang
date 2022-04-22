package dao

import (
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bot struct{}

// FindByID ...
func (p Bot) FindByID(ID primitive.ObjectID) (model.Bot, error) {
	var (
		botCol = database.BotCol()
		bot    model.Bot
	)

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := botCol.FindOne(context.Background(), filter).Decode(&bot); err != nil {
		fmt.Println(err)
		return bot, err
	}

	return bot, nil
}

// Create ...
func (p Bot) Create(bot model.Bot) error {
	var botCol = database.BotCol()

	// InsertOne
	if _, err := botCol.InsertOne(context.Background(), bot); err != nil {
		return err
	}

	return nil
}

// Update ...
func (p Bot) Update(ID primitive.ObjectID, bot model.Bot) error {
	var botCol = database.BotCol()

	// UpdateOne
	if _, err := botCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": bot}); err != nil {
		return err
	}

	return nil
}

// DeleteByID ...
func (p Bot) DeleteByID(ID primitive.ObjectID) error {
	var botCol = database.BotCol()

	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne ...
	if _, err := botCol.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

// DeleteAll ...
func (p Bot) DeleteAll() error {
	var botCol = database.BotCol()

	// DeleteMany ...
	if _, err := botCol.DeleteMany(context.Background(), bson.M{}); err != nil {
		return err
	}

	return nil
}

// GetList ...
func (p Bot) GetList(page, limit int) ([]model.Bot, error) {
	var (
		botCol = database.BotCol()
		bots   []model.Bot
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

	cursor, err := botCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return bots, err
	}

	if err = cursor.All(context.Background(), &bots); err != nil {
		return nil, err
	}

	fmt.Println(bots)

	return bots, nil
}

// CountAllBot ...
func (g Bot) CountAllBot() int {
	var (
		botCol = database.BotCol()
		ctx    = context.Background()
	)
	count, err := botCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0
	}
	return int(count)
}
