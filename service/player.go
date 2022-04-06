package service

import "card-game-golang/dto"

type Player struct{}

func (p Player) Register(player dto.Player) error {
	// call dao create player
	if err := playerDao.Create(player); err != nil {
		return err
	}

	return nil
}
