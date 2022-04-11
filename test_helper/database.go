package testhelper

import (
	"card-game-golang/dao"
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// ClearDB ...
func ClearDB() {
	database.PlayerCol().DeleteMany(context.Background(), bson.M{})
}

// CreateFakePlayer ...
func CreateFakePlayer() {
	var (
		playerDao = dao.Player{}
		bytes, _  = bcrypt.GenerateFromPassword([]byte("123456"), 14)

		// fakePlayer
		fakePlayer = model.Player{
			ID:        primitive.NewObjectID(),
			Name:      "fake",
			Email:     "fake@gmail.com",
			Password:  string(bytes),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	)

	err := playerDao.Create(fakePlayer)
	if err != nil {
		panic(err)
	}
}