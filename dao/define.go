package dao

import (
	"card-game-golang/dto"
	"card-game-golang/module/database"
)

// collections
var (
	playerCol = database.PlayerCol()
	adminCol  = database.AdminCol()
	botCol    = database.BotCol()
	gameCol   = database.GameCol()
	statsCol  = database.StatsCol()
)

type Service interface {
	Create(player dto.Player) error
}
