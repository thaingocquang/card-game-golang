package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// Game ...
	Game struct {
		ID         primitive.ObjectID `bson:"_id"`
		GameNo     int                `bson:"gameNo"`
		PlayerID   primitive.ObjectID `bson:"playerID"`
		BotID      primitive.ObjectID `bson:"botID"`
		WinnerID   primitive.ObjectID `bson:"winnerID"`
		PlayerHand Hand               `bson:"playerHand"`
		BotHand    Hand               `bson:"botHand"`
		BetValue   int                `bson:"betValue"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}

	// Card ...
	Card struct {
		Name string `bson:"name,omitempty"`
		Rank int    `bson:"rank,omitempty"`
		Suit int    `bson:"suit,omitempty"`
	}

	// Hand ...
	Hand struct {
		Cards   []Card `bson:"cards"`
		MaxCard Card   `bson:"maxCard"`
	}

	// DeckCard ...
	DeckCard []Card
)
