package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Auth ...
type Auth struct{}

// Register ...
func (a Auth) Register(player dto.Player) error {
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

	objID := primitive.NewObjectID()

	// player bson
	playerBSON := model.Player{
		ID:        objID,
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

	// dao create statistics
	if err := statsDao.Create(objID); err != nil {
		return err
	}

	return nil
}

// Login ...
func (a Auth) Login(player dto.PlayerLogin) (string, error) {
	// find player by email
	playerBSON, err := playerDao.FindByEmail(player.Email)
	if err != nil {
		return "", errors.New("email not existed in db")
	}

	// verify player password
	if err := bcrypt.CompareHashAndPassword([]byte(playerBSON.Password), []byte(player.Password)); err != nil {
		return "", errors.New("wrong password")
	}

	// jwt payload
	data := map[string]interface{}{
		"id": playerBSON.ID,
	}

	// GenerateUserToken ...
	token, err := util.GenerateUserToken(data)
	if err != nil {
		return "", errors.New("generate token failed")
	}

	// return JWT token
	return token, err
}

// AdminLogin ...
func (a Auth) AdminLogin(body dto.Admin) (string, error) {
	// find admin in db
	admin, err := adminDao.FindByUsername(body.Username)
	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.Password != body.Password {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id":    admin.ID,
		"admin": true,
	}

	// GenerateUserToken ...
	token, err := util.GenerateUserToken(data)
	if err != nil {
		return "", errors.New("generate token failed")
	}

	// return JWT token
	return token, err
}
