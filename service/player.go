package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Player struct{}

func (p Player) Register(player dto.Player) error {
	// check email existed
	isEmailExisted, err := playerDao.IsEmailExisted(player.Email)
	if err != nil {
		return errors.New("check mail existed failed")
	}

	// email existed
	if isEmailExisted == true {
		return errors.New("email already existed")
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(player.Password), 14)
	if err != nil {
		return err
	}

	// player bson
	playerBSON := model.Player{
		ID:        primitive.NewObjectID(),
		Name:      player.Name,
		Email:     player.Email,
		Password:  string(bytes),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// call dao create player
	if err := playerDao.Create(playerBSON); err != nil {
		return err
	}

	return nil
}
