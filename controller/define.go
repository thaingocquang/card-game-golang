package controller

import "card-game-golang/service"

var (
	authService   = service.Auth{}
	playerService = service.Player{}
	botService    = service.Bot{}
	gameService   = service.Game{}
)
