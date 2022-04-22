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
func (b Bot) GetByID(id string) (dto.BotJSON, error) {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// call dao get bot by id
	bot, err := botDao.FindByID(objID)
	if err != nil {
		return dto.BotJSON{}, err
	}

	botJSON := dto.BotJSON{
		ID:           bot.ID,
		Name:         bot.Name,
		TotalPoints:  bot.TotalPoints,
		RemainPoints: bot.RemainPoints,
		MinBet:       bot.MinBet,
		MaxBet:       bot.MaxBet,
	}

	// success
	return botJSON, err
}

// GetList ...
func (b Bot) GetList(page, limit int) ([]dto.BotJSON, int, error) {
	bots := make([]dto.BotJSON, 0)

	// call dao get list bot
	botBSONs, err := botDao.GetList(page, limit)
	if err != nil {
		return bots, 0, err
	}

	for _, bot := range botBSONs {
		botJSON := dto.BotJSON{
			ID:           bot.ID,
			Name:         bot.Name,
			TotalPoints:  bot.TotalPoints,
			RemainPoints: bot.RemainPoints,
			MinBet:       bot.MinBet,
			MaxBet:       bot.MaxBet,
		}
		bots = append(bots, botJSON)
	}

	totalDoc := botDao.CountAllBot()

	// success
	return bots, totalDoc, err
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
