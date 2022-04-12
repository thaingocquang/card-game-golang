package main

import (
	"card-game-golang/config"
	"card-game-golang/dao"
	"card-game-golang/module/database"
	"card-game-golang/route"
	"github.com/labstack/echo/v4"
)

// init ...
func init() {
	config.Init()
	database.Connect()
	dao.Admin{}.Create()
}

// main ...
func main() {
	//envVars ...
	envVars := config.GetEnv()

	//echo ...
	e := echo.New()

	//route
	route.Route(e)

	//start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
