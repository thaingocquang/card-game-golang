package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"time"
)

// Game ...
type Game struct{}

// PlayByBotID ...
func (g Game) PlayByBotID(gameVal dto.GameVal, botID string, myID string) (model.Game, error) {
	// game result
	gameBSON := model.Game{}

	// empty deck card
	var deckCard dto.DeckCard

	// init deck card
	deckCard.Init()

	// shuffle deck card
	deckCard.Shuffle()

	// deal card
	var botHand, playerHand dto.Hand
	botHand.Cards = deckCard[0:3]
	playerHand.Cards = deckCard[3:6]

	// find max card
	botHand.InitMaxCard()
	playerHand.InitMaxCard()

	// convert id to objectID
	myObjID, _ := primitive.ObjectIDFromHex(myID)
	botObjID, _ := primitive.ObjectIDFromHex(botID)

	myStatsBSON, err := statsDao.FindByPlayerID(myObjID)
	if err != nil {
		return gameBSON, err
	}
	botBSON, err := botDao.FindByID(botObjID)
	if err != nil {
		return gameBSON, err
	}

	// init BSON doc for update
	myStatsUpdateBSON := model.StatsUpdate{}
	botUpdateBSON := model.Bot{}

	// increase player totalGame by 1
	myStatsUpdateBSON.TotalGame = myStatsBSON.TotalGame + 1

	// compare hand
	if playerHand.CompareHandIsHigher(botHand) {
		// if player win
		gameBSON.WinnerID = myObjID

		// increase player winGame by 1
		myStatsUpdateBSON.WinGame = myStatsBSON.WinGame + 1

		// calculate player winRate
		myStatsUpdateBSON.WinRate = float32(myStatsUpdateBSON.WinGame) / float32(myStatsUpdateBSON.TotalGame)

		// add bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point + gameVal.BetValue

		// minus bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints - gameVal.BetValue
		if botUpdateBSON.RemainPoints < 0 {
			botUpdateBSON.RemainPoints = 0
		}
	} else {
		// if bot win
		gameBSON.WinnerID = botObjID

		// add bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints + gameVal.BetValue

		// minus bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point - gameVal.BetValue
		if myStatsUpdateBSON.Point < 0 {
			myStatsUpdateBSON.Point = 0
		}
	}

	// update my stats
	if err = statsDao.UpdateByPlayerID(myObjID, myStatsUpdateBSON); err != nil {
		return gameBSON, err
	}

	// update bot
	if err = botDao.Update(botObjID, botUpdateBSON); err != nil {
		return gameBSON, err
	}

	// record game result
	gameBSON.ID = primitive.NewObjectID()
	gameBSON.GameNo = gameDao.CountAllGame()
	gameBSON.PlayerID = myObjID
	gameBSON.BotID = botObjID
	gameBSON.PlayerHand = playerHand
	gameBSON.BotHand = botHand
	gameBSON.BetValue = gameVal.BetValue
	gameBSON.CreatedAt = time.Now()
	gameBSON.UpdatedAt = time.Now()

	// create game
	if err := gameDao.Create(gameBSON); err != nil {
		return model.Game{}, err
	}

	return gameBSON, nil

}

func (g Game) PlayRandom(gameVal dto.GameVal, myID string) (model.Game, error) {
	gameBSON := model.Game{}

	// get all bot
	bots, err := botDao.GetList(0, 0)
	if err != nil {
		return gameBSON, err
	}

	// filter bot have totalPoint > betValue & in range (minBet, maxBet)
	validBots := make([]model.Bot, 0)
	for _, bot := range bots {
		if gameVal.BetValue >= bot.MinBet && gameVal.BetValue <= bot.MaxBet {
			if bot.TotalPoints >= gameVal.BetValue {
				validBots = append(validBots, bot)
			}
		}
	}

	if len(validBots) == 0 {
		return gameBSON, errors.New("no bot satisfy betValue")
	}

	// random bot in list
	botBSON := model.Bot{}

	if len(validBots) == 1 {
		botBSON = validBots[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		botBSON = validBots[rand.Intn(len(validBots)-1)]
	}

	// empty deck card
	var deckCard dto.DeckCard

	// init deck card
	deckCard.Init()

	// shuffle deck card
	deckCard.Shuffle()

	// deal card
	var botHand, playerHand dto.Hand
	botHand.Cards = deckCard[0:3]
	playerHand.Cards = deckCard[3:6]

	// find max card
	botHand.InitMaxCard()
	playerHand.InitMaxCard()

	// convert id to objectID
	myObjID, _ := primitive.ObjectIDFromHex(myID)

	myStatsBSON, err := statsDao.FindByPlayerID(myObjID)
	if err != nil {
		return gameBSON, err
	}
	//botBSON, err := botDao.FindByID(botObjID)
	//if err != nil {
	//	return gameBSON, err
	//}

	// init BSON doc for update
	myStatsUpdateBSON := model.StatsUpdate{}
	botUpdateBSON := model.Bot{}

	// increase player totalGame by 1
	myStatsUpdateBSON.TotalGame = myStatsBSON.TotalGame + 1

	// compare hand
	if playerHand.CompareHandIsHigher(botHand) {
		// if player win
		gameBSON.WinnerID = myObjID

		// increase player winGame by 1
		myStatsUpdateBSON.WinGame = myStatsBSON.WinGame + 1

		// calculate player winRate
		myStatsUpdateBSON.WinRate = float32(myStatsUpdateBSON.WinGame) / float32(myStatsUpdateBSON.TotalGame)

		// add bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point + gameVal.BetValue

		// minus bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints - gameVal.BetValue
		if botUpdateBSON.RemainPoints < 0 {
			botUpdateBSON.RemainPoints = 0
		}
	} else {
		// if bot win
		gameBSON.WinnerID = botBSON.ID

		// add bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints + gameVal.BetValue

		// minus bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point - gameVal.BetValue
		if myStatsUpdateBSON.Point < 0 {
			myStatsUpdateBSON.Point = 0
		}
	}

	// update my stats
	if err = statsDao.UpdateByPlayerID(myObjID, myStatsUpdateBSON); err != nil {
		return gameBSON, err
	}

	// update bot
	if err = botDao.Update(botBSON.ID, botUpdateBSON); err != nil {
		return gameBSON, err
	}

	// record game result
	gameBSON.ID = primitive.NewObjectID()
	gameBSON.GameNo = gameDao.CountAllGame()
	gameBSON.PlayerID = myObjID
	gameBSON.BotID = botBSON.ID
	gameBSON.PlayerHand = playerHand
	gameBSON.BotHand = botHand
	gameBSON.BetValue = gameVal.BetValue
	gameBSON.CreatedAt = time.Now()
	gameBSON.UpdatedAt = time.Now()

	// create game
	if err := gameDao.Create(gameBSON); err != nil {
		return model.Game{}, err
	}

	return gameBSON, nil
}

// GetList ...
func (g Game) GetList(page, limit int) ([]model.Game, int, error) {
	// call dao get list bot
	games, err := gameDao.GetList(page, limit)
	if err != nil {
		return games, 0, err
	}

	totalDoc := gameDao.CountAllGame()

	// success
	return games, totalDoc, err
}
