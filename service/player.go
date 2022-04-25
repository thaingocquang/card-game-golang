package service

import (
	"card-game-golang/dto"
	"card-game-golang/model"
	"card-game-golang/util"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Player ...
type Player struct{}

// GetByID ...
func (p Player) GetByID(id string) (dto.MyProfile, error) {
	var profile dto.MyProfile

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

	profile = dto.MyProfile{
		ID:        player.ID,
		Name:      player.Name,
		Email:     player.Email,
		Password:  player.Password,
		Point:     stats.Point,
		TotalGame: stats.TotalGame,
		WinGame:   stats.WinGame,
		WinRate:   stats.WinRate,
	}

	return profile, nil
}

// UpdateProfile ...
func (p Player) UpdateProfile(ID string, update dto.ProfileUpdate) error {
	// get objectID from string
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	// update point to stats
	statsUpdateBSON := model.StatsUpdate{Point: update.Point}

	if err := statsDao.UpdateByPlayerID(objID, statsUpdateBSON); err != nil {
		return err
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(update.Password), 14)
	if err != nil {
		return err
	}

	// update player to playerDao
	playerUpdateBSON := model.Player{
		Name:     update.Name,
		Email:    update.Email,
		Password: string(bytes),
	}
	if err := playerDao.Update(objID, playerUpdateBSON); err != nil {
		return err
	}

	return nil
}

// Update ...
func (p Player) Update(ID string, update dto.PlayerUpdate) error {
	// get objectID from string
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	player, err := playerDao.FindByID(objID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(update.Password)); err != nil {
		return errors.New("wrong password, please input correct password again")
	}

	if update.NewPassword != "" {
		// hash player password
		bytes, err := bcrypt.GenerateFromPassword([]byte(update.NewPassword), 14)
		if err != nil {
			return err
		}

		// update player to playerDao
		playerUpdateBSON := model.Player{
			Name:     update.Name,
			Email:    update.Email,
			Password: string(bytes),
		}

		// call dao update player
		if err := playerDao.Update(objID, playerUpdateBSON); err != nil {
			return err
		}

	} else {
		// update player to playerDao
		playerUpdateBSON := model.Player{
			Name:  update.Name,
			Email: update.Email,
		}

		// call dao update player
		if err := playerDao.Update(objID, playerUpdateBSON); err != nil {
			return err
		}
	}

	return nil
}

//// GetList ...
//func (p Player) GetList(page, limit int) ([]dto.Profile, int, error) {
//	profiles := make([]dto.Profile, 0)
//
//	// get player
//	players, err := playerDao.GetList(page, limit)
//	if err != nil {
//		return profiles, 0, err
//	}
//
//	// get statistics
//	stats, err := statsDao.GetList(page, limit)
//	if err != nil {
//		return profiles, 0, err
//	}
//
//	if len(players) == len(stats) {
//		for i, player := range players {
//			profile := dto.Profile{
//				ID:        player.ID,
//				Name:      player.Name,
//				Email:     player.Email,
//				Point:     stats[i].Point,
//				TotalGame: stats[i].TotalGame,
//				WinGame:   stats[i].WinGame,
//				WinRate:   stats[i].WinRate,
//			}
//			profiles = append(profiles, profile)
//		}
//	}
//
//	totalDoc := playerDao.CountAllPlayer()
//
//	return profiles, totalDoc, nil
//}

// GetListProfile ...
func (p Player) GetListProfile(paging *util.Paging) ([]dto.Profile, error) {
	profilesJSON := make([]dto.Profile, 0)

	// get profile
	profilesBSON, err := playerDao.GetListProfile(paging)
	if err != nil {
		return nil, err
	}

	for _, profileBSON := range profilesBSON {
		profileJSON := dto.Profile{
			ID:        profileBSON.ID,
			Name:      profileBSON.Name,
			Email:     profileBSON.Email,
			Point:     profileBSON.Stat.Point,
			TotalGame: profileBSON.Stat.TotalGame,
			WinGame:   profileBSON.Stat.WinGame,
			WinRate:   profileBSON.Stat.WinRate,
		}
		profilesJSON = append(profilesJSON, profileJSON)
	}

	return profilesJSON, nil
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

//// Test ...
//func (p Player) Test() error {
//	// delete in player collection
//	if err := playerDao.GetListAggregate(); err != nil {
//		return err
//	}
//
//	return nil
//}
