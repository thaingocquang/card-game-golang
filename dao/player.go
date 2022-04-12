package dao

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/module/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct{}

// Create ...
func (p Player) Create(player model.Player) error {
	var playerCol = database.PlayerCol()

	// InsertOne
	if _, err := playerCol.InsertOne(context.Background(), player); err != nil {
		return err
	}

	return nil
}

// FindByID ...
func (p Player) FindByID(ID primitive.ObjectID) (model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		player    model.Player
	)

	// filter
	filter := bson.M{"_id": ID}

	// FindOne
	if err := playerCol.FindOne(context.Background(), filter).Decode(&player); err != nil {
		return model.Player{}, err
	}

	return player, nil
}

// FindByEmail ...
func (p Player) FindByEmail(email string) (model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		player    model.Player
	)

	// filter
	filter := bson.M{"email": email}

	// FindOne
	if err := playerCol.FindOne(context.Background(), filter).Decode(&player); err != nil {
		return model.Player{}, err
	}

	return player, nil
}

// Update ...
func (p Player) Update(ID primitive.ObjectID, player dto.Player) error {
	var playerCol = database.PlayerCol()

	update := model.Player{Name: player.Name, Email: player.Email, Password: player.Password}

	// UpdateOne
	if _, err := playerCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": update}); err != nil {
		return err
	}

	return nil
}

// UpdateProfile ...
func (p Player) UpdateProfile(ID primitive.ObjectID, player dto.PlayerUpdate) error {
	var playerCol = database.PlayerCol()

	update := model.Player{Name: player.Name, Email: player.Email}

	// UpdateOne
	if _, err := playerCol.UpdateOne(context.Background(), bson.M{"_id": ID}, bson.M{"$set": update}); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (p Player) Delete(ID primitive.ObjectID) error {
	var playerCol = database.PlayerCol()

	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne ...
	if _, err := playerCol.DeleteOne(context.Background(), filter); err != nil {
		return err
	}

	return nil
}

// GetList ...
func (p Player) GetList(page, limit int) ([]model.Player, error) {
	var (
		playerCol = database.PlayerCol()
		players   []model.Player
	)

	// options
	opts := new(options.FindOptions)

	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}

	cursor, err := playerCol.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &players); err != nil {
		return nil, err
	}

	return players, nil
}

func (p Player) IsEmailExisted(email string) (bool, error) {
	var playerCol = database.PlayerCol()

	// filter
	filter := bson.M{"email": email}

	//
	count, err := playerCol.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
