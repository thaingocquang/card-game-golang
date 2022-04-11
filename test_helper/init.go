package testhelper

import (
	"card-game-golang/config"
	"card-game-golang/module/database"
	"card-game-golang/route"
	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	// Init initialize app's config
	config.Init()

	database.Connect()

	// ClearDB ...
	ClearDB()

	// new test server
	e := echo.New()

	// route
	route.Route(e)

	return e
}
