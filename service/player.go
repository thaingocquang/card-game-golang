package service

import (
	"card-game-golang/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player ...
type Player struct{}

// GetByID ...
func (p Player) GetByID(id string) (dto.Profile, error) {
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
		ID:        player.ID,
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

// GetList ...
func (p Player) GetList(page, limit int) ([]dto.Profile, error) {
	profiles := make([]dto.Profile, 0)

	// get player
	players, err := playerDao.GetList(page, limit)
	if err != nil {
		return profiles, err
	}

	// get statistics
	stats, err := statsDao.GetList(page, limit)
	if err != nil {
		return profiles, err
	}

	if len(players) == len(stats) {
		for i, player := range players {
			profile := dto.Profile{
				ID:        player.ID,
				Name:      player.Name,
				Email:     player.Email,
				Point:     stats[i].Point,
				TotalGame: stats[i].TotalGame,
				WinGame:   stats[i].WinGame,
				WinRate:   stats[i].WinRate,
			}
			profiles = append(profiles, profile)
		}
	}

	return profiles, nil
}

// DeleteByID ...
func (p Player) DeleteByID(ID string) error {
	// convert id to objectID
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	// delete in player collection
	if err = playerDao.DeleteByID(objID); err != nil {
		return err
	}

	// delete in stats collection
	if err = statsDao.DeleteByID(objID); err != nil {
		return err
	}

	return nil
}

// DeleteAll ...
func (p Player) DeleteAll() error {
	// delete in player collection
	if err := playerDao.DeleteAll(); err != nil {
		return err
	}

	// delete in stats collection
	if err := statsDao.DeleteAll(); err != nil {
		return err
	}

	return nil
}
