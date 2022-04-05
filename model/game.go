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
		CreatedAt  time.Time          `bson:"created_at"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}
)
