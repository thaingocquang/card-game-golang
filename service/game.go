package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/util"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strconv"
	"time"
)

// Game ...
type Game struct{}

// InitGame ...
func initGame() (dto.Hand, dto.Hand) {
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

	return botHand, playerHand
}

func recordGame(botId primitive.ObjectID, playerID primitive.ObjectID, botHand dto.Hand, playerHand dto.Hand, gameVal dto.GameVal) (dto.GameJSON, error) {
	// game result
	gameBSON := model.Game{}
	gameJSON := dto.GameJSON{}

	// get player
	player, err := playerDao.FindByID(playerID)
	if err != nil {
		return gameJSON, err
	}

	myStatsBSON, err := statsDao.FindByPlayerID(playerID)
	if err != nil {
		return gameJSON, err
	}
	botBSON, err := botDao.FindByID(botId)
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
		gameBSON.WinnerID = playerID
		gameBSON.WinnerName = player.Name
		gameJSON.WinnerID = playerID
		gameJSON.WinnerName = player.Name

		// increase player winGame by 1
		myStatsUpdateBSON.WinGame = myStatsBSON.WinGame + 1

		// calculate player winRate
		myStatsUpdateBSON.WinRate = float32(myStatsUpdateBSON.WinGame) / float32(myStatsUpdateBSON.TotalGame)

		// add bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point + gameVal.BetValue

		// minus bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints - gameVal.BetValue
		if botUpdateBSON.RemainPoints <= 0 {
			botUpdateBSON.RemainPoints = 0
			fmt.Println("hello")
		}
	} else {
		// if bot win
		gameBSON.WinnerID = botId
		gameBSON.WinnerName = botBSON.Name
		gameJSON.WinnerID = botId
		gameJSON.WinnerName = botBSON.Name

		// add bet value to bot
		botUpdateBSON.RemainPoints = botBSON.RemainPoints + gameVal.BetValue

		// minus bet value to player
		myStatsUpdateBSON.Point = myStatsBSON.Point - gameVal.BetValue
		if myStatsUpdateBSON.Point <= 0 {
			myStatsUpdateBSON.Point = 0
		}
	}

	fmt.Println(botUpdateBSON)

	// update my stats
	if err = statsDao.UpdateByPlayerID(playerID, myStatsUpdateBSON); err != nil {
		return gameJSON, err
	}

	// update bot
	if err = botDao.Update(botId, botUpdateBSON); err != nil {
		return gameJSON, err
	}

	// record game result
	gameID := primitive.NewObjectID()

	gameBSON.ID = gameID
	gameBSON.GameNo = gameDao.CountAllGame()
	gameBSON.PlayerID = playerID
	gameBSON.BotID = botId
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
	gameJSON.ID = gameID
	gameJSON.GameNo = gameDao.CountAllGame()
	gameJSON.PlayerID = playerID
	gameJSON.BotID = botId
	gameJSON.PlayerHand = playerHand
	gameJSON.BotHand = botHand
	gameJSON.BetValue = gameVal.BetValue
	gameJSON.CreatedAt = time.Now()
	gameJSON.UpdatedAt = time.Now()

	return gameJSON, nil
}

// PlayByBotID ...
func (g Game) PlayByBotID(gameVal dto.GameVal, botID string, myID string) (dto.GameJSON, error) {
	// convert id to objectID
	myObjID, _ := primitive.ObjectIDFromHex(myID)
	botObjID, _ := primitive.ObjectIDFromHex(botID)

	// check bet value
	botBSON, err := botDao.FindByID(botObjID)
	if err != nil {
		return dto.GameJSON{}, err
	}
	if gameVal.BetValue > botBSON.MaxBet || gameVal.BetValue < botBSON.MinBet {
		return dto.GameJSON{}, errors.New("bet value not satisfy")
	}

	// init game
	botHand, playerHand := initGame()

	// record game
	gameJSON, err := recordGame(botObjID, myObjID, botHand, playerHand, gameVal)
	if err != nil {
		return gameJSON, err
	}

	return gameJSON, nil
}

// getListBotSatisfyBotVal ...
func getListBotSatisfyBetVal(gameVal dto.GameVal) ([]model.Bot, error) {
	// filter bot have totalPoint > betValue & in range (minBet, maxBet)
	validBots := make([]model.Bot, 0)

	//get all bot
	bots, err := botDao.GetAll()
	if err != nil {
		return validBots, err
	}

	for _, bot := range bots {
		if gameVal.BetValue >= bot.MinBet && gameVal.BetValue <= bot.MaxBet {
			if bot.RemainPoints >= gameVal.BetValue {
				validBots = append(validBots, bot)
			}
		}
	}

	return validBots, nil
}

// randomBotInList ...
func randomBotInList(validBots []model.Bot) model.Bot {
	// random bot in list
	botBSON := model.Bot{}

	if len(validBots) == 1 {
		botBSON = validBots[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		botBSON = validBots[rand.Intn(len(validBots)-1)]
	}

	return botBSON
}

// PlayRandom ...
func (g Game) PlayRandom(gameVal dto.GameVal, myID string) (dto.GameJSON, error) {
	gameJSON := dto.GameJSON{}

	// convert id to objectID
	myObjID, _ := primitive.ObjectIDFromHex(myID)

	myStats, err := statsDao.FindByPlayerID(myObjID)
	if err != nil {
		return dto.GameJSON{}, err
	}

	validBots, err := getListBotSatisfyBetVal(gameVal)
	if err != nil {
		return gameJSON, err
	}

	if len(validBots) == 0 {
		botBSONs, err := botDao.GetAll()
		if err != nil {
			return dto.GameJSON{}, err
		}
		minBetInList := 999999
		maxBetInList := 0
		//for _, v := range botBSONs {
		//	if minBetInList > v.MinBet {
		//		minBetInList = v.MinBet
		//	}
		//	if maxBetInList < v.MaxBet {
		//		maxBetInList = v.MaxBet
		//	}
		//}
		for _, bot := range botBSONs {
			if bot.RemainPoints >= bot.MaxBet {
				if myStats.Point >= bot.MaxBet {
					if maxBetInList < bot.MaxBet {
						maxBetInList = bot.MaxBet
					}
				}
			} else {
				if myStats.Point >= bot.MinBet {
					if maxBetInList < bot.RemainPoints {
						maxBetInList = bot.RemainPoints
					}
				}
			}
			if bot.RemainPoints >= bot.MinBet && myStats.Point >= bot.MinBet {
				if minBetInList > bot.MinBet {
					minBetInList = bot.MinBet
				}
			}
		}

		if minBetInList != 999999 || maxBetInList != 0 {
			return gameJSON, errors.New("Không có bot nào thỏa mãn giá trị cược, bạn có thể cược trong khoảng " + strconv.Itoa(minBetInList) + " đến " + strconv.Itoa(maxBetInList))
		} else {
			return gameJSON, errors.New("Không có bot nào thỏa mãn giá trị cược.")
		}

	}

	botBSON := randomBotInList(validBots)

	//// pick bot which have RemainPoints > 0
	//for {
	//	if botBSON.RemainPoints > 0 || botBSON.RemainPoints < gameVal.BetValue {
	//		break
	//	} else {
	//		botBSON = randomBotInList(validBots)
	//	}
	//}

	// init game
	botHand, playerHand := initGame()

	gameJSON, err = recordGame(botBSON.ID, myObjID, botHand, playerHand, gameVal)
	if err != nil {
		return gameJSON, err
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
		botHandJSON := dto.Hand{}
		playerHandJSON := dto.Hand{}
		for _, v := range game.BotHand.Cards {
			card := dto.Card{
				Name: v.Name,
				Rank: v.Rank,
				Suit: v.Suit,
			}
			botHandJSON.Cards = append(botHandJSON.Cards, card)
			//botHandJSON.Cards[i].Rank = v.Rank
			//botHandJSON.Cards[i].Suit = v.Suit
			//botHandJSON.Cards[i].Name = v.Name
		}
		botHandJSON.MaxCard.Name = game.BotHand.MaxCard.Name
		botHandJSON.MaxCard.Suit = game.BotHand.MaxCard.Suit
		botHandJSON.MaxCard.Rank = game.BotHand.MaxCard.Rank

		for _, v := range game.PlayerHand.Cards {
			card := dto.Card{
				Name: v.Name,
				Rank: v.Rank,
				Suit: v.Suit,
			}
			playerHandJSON.Cards = append(playerHandJSON.Cards, card)
			//playerHandJSON.Cards[i].Rank = v.Rank
			//playerHandJSON.Cards[i].Suit = v.Suit
			//playerHandJSON.Cards[i].Name = v.Name
		}
		playerHandJSON.MaxCard.Name = game.PlayerHand.MaxCard.Name
		playerHandJSON.MaxCard.Suit = game.PlayerHand.MaxCard.Suit
		playerHandJSON.MaxCard.Rank = game.PlayerHand.MaxCard.Rank

		gameJSON := dto.GameJSON{
			ID:         game.ID,
			GameNo:     game.GameNo,
			PlayerID:   game.PlayerID,
			BotID:      game.BotID,
			WinnerID:   game.WinnerID,
			WinnerName: game.WinnerName,
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
