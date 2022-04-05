package dao

import "card-game-golang/module/database"

var (
	playerCol = database.PlayerCol()
	adminCol  = database.AdminCol()
	botCol    = database.BotCol()
	gameCol   = database.GameCol()
	statsCol  = database.StatsCol()
)
