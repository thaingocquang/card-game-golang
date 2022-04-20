package route

import (
	"card-game-golang/controller"
	"card-game-golang/validations"
)

var (
	// define controller
	botCtrl    = controller.Bot{}
	authCtrl   = controller.Auth{}
	playerCtrl = controller.Player{}
	gameCtrl   = controller.Game{}

	// define validation
	playerVal = validations.Player{}
	val       = validations.Validator{}
	adminVal  = validations.Admin{}
	botVal    = validations.Bot{}
	gameVal   = validations.Game{}
)
