package dao

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
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
		return model.Bot{}, err
	}

	return bot, nil
}

// Create ...
func (p Bot) Create(bot dto.Bot) error {
	var botCol = database.BotCol()

	// InsertOne
	if _, err := botCol.InsertOne(context.Background(), bot); err != nil {
		return err
	}

	return nil
}

// Update ...
func (p Bot) Update(ID primitive.ObjectID, bot dto.Bot) error {
	var botCol = database.BotCol()

	update := model.Bot{
		Name:         bot.Name,
		TotalPoints:  bot.TotalPoints,
		RemainPoints: bot.RemainPoints,
		MinBet:       bot.MinBet,
		MaxBet:       bot.MaxBet,
	}

	// UpdateOne
	if _, err := botCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": update}); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (p Bot) Delete(ID primitive.ObjectID) error {
	var botCol = database.BotCol()

	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne ...
	if _, err := botCol.DeleteOne(context.Background(), filter); err != nil {
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
		return nil, err
	}

	if err = cursor.All(context.Background(), &bots); err != nil {
		return nil, err
	}

	return bots, nil
}
