package main

import (
	"card-game-golang/config"
	"card-game-golang/dao"
	_ "card-game-golang/docs"
	"card-game-golang/module/database"
	"card-game-golang/route"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// init ...
func init() {
	config.Init()
	database.Connect()
	dao.Admin{}.Create()
}

// @title Card Game API
// @version 1.0
// @description This is a Card Game server.
// @license.name Apache 2.0
// @host localhost:1323
// @BasePath /
func main() {
	//envVars ...
	envVars := config.GetEnv()

	//echo ...
	e := echo.New()

	//route
	route.Route(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
