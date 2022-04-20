package service

import "card-game-golang/dao"

var (
	playerDao = dao.Player{}
	statsDao  = dao.Stats{}
	adminDao  = dao.Admin{}
	botDao    = dao.Bot{}
	gameDao   = dao.Game{}
)
