package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"time"
)

// Game ...
type Game struct{}

// PlayByBotID ...
func (g Game) PlayByBotID(gameVal dto.GameVal, botID string, myID string) (dto.GameJSON, error) {
	// game result
	gameBSON := model.Game{}
	gameJSON := dto.GameJSON{}

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
		return gameJSON, err
	}
	botBSON, err := botDao.FindByID(botObjID)
	if err != nil {
		return gameJSON, err
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
		return gameJSON, err
	}

	// update bot
	if err = botDao.Update(botObjID, botUpdateBSON); err != nil {
		return gameJSON, err
	}

	// record game result
	gameBSON.ID = primitive.NewObjectID()
	gameBSON.GameNo = gameDao.CountAllGame()
	gameBSON.PlayerID = myObjID
	gameBSON.BotID = botObjID
	gameBSON.PlayerHand = playerHand.ConvertToBSON()
	gameBSON.BotHand = botHand.ConvertToBSON()
	gameBSON.BetValue = gameVal.BetValue
	gameBSON.CreatedAt = time.Now()
	gameBSON.UpdatedAt = time.Now()

	// create game
	if err := gameDao.Create(gameBSON); err != nil {
		return gameJSON, err
	}

	// gameJSON result
	gameJSON.ID = primitive.NewObjectID()
	gameJSON.GameNo = gameDao.CountAllGame()
	gameJSON.PlayerID = myObjID
	gameJSON.BotID = botObjID
	gameJSON.PlayerHand = playerHand
	gameJSON.BotHand = botHand
	gameJSON.BetValue = gameVal.BetValue
	gameJSON.CreatedAt = time.Now()
	gameJSON.UpdatedAt = time.Now()

	return gameJSON, nil

}

func (g Game) PlayRandom(gameVal dto.GameVal, myID string) (dto.GameJSON, error) {
	gameBSON := model.Game{}
	gameJSON := dto.GameJSON{}

	// get all bot
	//bots, err := botDao.GetList(0, 0)
	//if err != nil {
	//	return gameJSON, err
	//}

	// delete
	var bots []model.Bot

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
		return gameJSON, errors.New("no bot satisfy betValue")
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
		return gameJSON, err
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
		return gameJSON, err
	}

	// update bot
	if err = botDao.Update(botBSON.ID, botUpdateBSON); err != nil {
		return gameJSON, err
	}

	// record game result
	gameBSON.ID = primitive.NewObjectID()
	gameBSON.GameNo = gameDao.CountAllGame()
	gameBSON.PlayerID = myObjID
	gameBSON.BotID = botBSON.ID
	gameBSON.PlayerHand = playerHand.ConvertToBSON()
	gameBSON.BotHand = botHand.ConvertToBSON()
	gameBSON.BetValue = gameVal.BetValue
	gameBSON.CreatedAt = time.Now()
	gameBSON.UpdatedAt = time.Now()

	// create game
	if err := gameDao.Create(gameBSON); err != nil {
		return gameJSON, err
	}

	gameJSON = dto.GameJSON{
		ID:         gameBSON.ID,
		GameNo:     gameBSON.GameNo,
		PlayerID:   gameBSON.PlayerID,
		BotID:      gameBSON.BotID,
		WinnerID:   gameBSON.WinnerID,
		PlayerHand: playerHand,
		BotHand:    botHand,
		BetValue:   gameBSON.BetValue,
		CreatedAt:  gameBSON.CreatedAt,
		UpdatedAt:  gameBSON.UpdatedAt,
	}

	return gameJSON, nil
}

// GetList ...
func (g Game) GetList(paging *util.Paging) ([]dto.GameJSON, error) {
	gameJSONs := make([]dto.GameJSON, 0)

	// call dao get list bot
	gameBSONs, err := gameDao.GetList(paging)
	if err != nil {
		return gameJSONs, err
	}

	for _, game := range gameBSONs {
		var playerHandJSON dto.Hand
		for _, v := range game.PlayerHand.Cards {
			playerHandJSON.Cards = append(playerHandJSON.Cards, dto.Card{
				Name: v.Name,
				Rank: v.Rank,
				Suit: v.Suit,
			})
		}
		playerHandJSON.MaxCard = dto.Card{
			Name: game.PlayerHand.MaxCard.Name,
			Rank: game.PlayerHand.MaxCard.Rank,
			Suit: game.PlayerHand.MaxCard.Suit,
		}

		var botHandJSON dto.Hand
		for _, v := range game.PlayerHand.Cards {
			botHandJSON.Cards = append(botHandJSON.Cards, dto.Card{
				Name: v.Name,
				Rank: v.Rank,
				Suit: v.Suit,
			})
		}
		botHandJSON.MaxCard = dto.Card{
			Name: game.PlayerHand.MaxCard.Name,
			Rank: game.PlayerHand.MaxCard.Rank,
			Suit: game.PlayerHand.MaxCard.Suit,
		}

		gameJSON := dto.GameJSON{
			ID:         game.ID,
			GameNo:     game.GameNo,
			PlayerID:   game.PlayerID,
			BotID:      game.BotID,
			WinnerID:   game.WinnerID,
			PlayerHand: playerHandJSON,
			BotHand:    botHandJSON,
			BetValue:   game.BetValue,
			CreatedAt:  game.CreatedAt,
			UpdatedAt:  game.UpdatedAt,
		}
		gameJSONs = append(gameJSONs, gameJSON)
	}

	// success
	return gameJSONs, err
}

// Recent ...
func (g Game) Recent(ID string) ([]dto.GameJSON, error) {
	gameJSONs := make([]dto.GameJSON, 0)

	objID, err := primitive.ObjectIDFromHex(ID)

	// call dao get list bot
	gameBSONs, err := gameDao.RecentByPlayerID(objID)
	if err != nil {
		return gameJSONs, err
	}

	for _, game := range gameBSONs {
		gameJSON := dto.GameJSON{
			ID:         game.ID,
			GameNo:     game.GameNo,
			PlayerID:   game.PlayerID,
			BotID:      game.BotID,
			WinnerID:   game.WinnerID,
			PlayerHand: dto.Hand{},
			BotHand:    dto.Hand{},
			BetValue:   game.BetValue,
			CreatedAt:  game.CreatedAt,
			UpdatedAt:  game.UpdatedAt,
		}
		gameJSONs = append(gameJSONs, gameJSON)
	}

	// success
	return gameJSONs, err
}
