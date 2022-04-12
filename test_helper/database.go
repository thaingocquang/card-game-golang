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

var (
	PlayerIDString = "5f24d45125ea51bc57a8285a"
	PlayerObjID, _ = primitive.ObjectIDFromHex(PlayerIDString)
)

// ClearDB ...
func ClearDB() {
	database.PlayerCol().DeleteMany(context.Background(), bson.M{})
}

// CreateFakePlayer ...
func CreateFakePlayer() {
	var (
		playerDao = dao.Player{}
		statsDao  = dao.Stats{}

		// player bson
		playerBSON = model.Player{
			ID:        PlayerObjID,
			Name:      "fake",
			Email:     "fake@gmail.com",
			Password:  "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	)
	// check email existed
	isEmailExisted, err := playerDao.IsEmailExisted(playerBSON.Email)
	if err != nil {
		panic(err)
	}

	// email existed
	if isEmailExisted == true {
		panic(err)
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(playerBSON.Password), 14)
	if err != nil {
		panic(err)
	}

	playerBSON.Password = string(bytes)

	// call dao create player
	if err := playerDao.Create(playerBSON); err != nil {
		panic(err)
	}

	// dao create statistics
	if err := statsDao.Create(PlayerObjID); err != nil {
		panic(err)
	}
}
