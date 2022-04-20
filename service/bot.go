package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Bot ...
type Bot struct{}

// Create ...
func (b Bot) Create(bot dto.Bot) error {
	// convert to BSON
	botBSON := model.Bot{
		ID:           primitive.NewObjectID(),
		Name:         bot.Name,
		TotalPoints:  bot.TotalPoints,
		RemainPoints: bot.TotalPoints,
		MinBet:       bot.MinBet,
		MaxBet:       bot.MaxBet,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// call dao create bot
	err := botDao.Create(botBSON)
	if err != nil {
		return err
	}

	// success
	return nil
}

// GetByID ...
func (b Bot) GetByID(id string) (model.Bot, error) {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// call dao get bot by id
	bot, err := botDao.FindByID(objID)
	if err != nil {
		return bot, err
	}

	// success
	return bot, err
}

// GetList ...
func (b Bot) GetList(page, limit int) ([]model.Bot, error) {
	// call dao get list bot
	bots, err := botDao.GetList(page, limit)
	if err != nil {
		return bots, err
	}

	// success
	return bots, err
}

// UpdateByID ...
func (b Bot) UpdateByID(id string, bot dto.Bot) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// convert to BSON
	botUpdateBSON := model.Bot{
		Name:         bot.Name,
		TotalPoints:  bot.TotalPoints,
		RemainPoints: bot.RemainPoints,
		MinBet:       bot.MinBet,
		MaxBet:       bot.MaxBet,
		UpdatedAt:    time.Now(),
	}

	// call dao update bot by id
	err := botDao.Update(objID, botUpdateBSON)
	if err != nil {
		return err
	}

	// success
	return nil
}

// DeleteByID ...
func (b Bot) DeleteByID(id string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// call dao delete bot by id
	err := botDao.DeleteByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}

// DeleteAll ...
func (b Bot) DeleteAll() error {
	// call dao delete all bot
	err := botDao.DeleteAll()
	if err != nil {
		return err
	}

	// success
	return nil
}
