package service

import (
	"card-game-golang/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player ...
type Player struct{}

// MyProfile ...
func (p Player) MyProfile(id string) (dto.Profile, error) {
	var profile dto.Profile

	// get objectID from string
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return profile, err
	}

	// get player
	player, err := playerDao.FindByID(objID)
	if err != nil {
		return profile, err
	}

	// get statistics
	stats, err := statsDao.FindByPlayerID(objID)
	if err != nil {
		return profile, err
	}

	profile = dto.Profile{
		Name:      player.Name,
		Email:     player.Email,
		Point:     stats.Point,
		TotalGame: stats.TotalGame,
		WinGame:   stats.WinGame,
		WinRate:   stats.WinRate,
	}

	return profile, nil
}

// UpdateProfile ...
func (p Player) UpdateProfile(ID string, update dto.PlayerUpdate) error {
	// get objectID from string
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	// UpdateProfile
	if err := playerDao.UpdateProfile(objID, update); err != nil {
		return err
	}

	return nil
}
