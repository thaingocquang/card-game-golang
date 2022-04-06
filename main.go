package main

import (
	"card-game-golang/config"
	"card-game-golang/module/database"
	"card-game-golang/route"
	"fmt"
	"github.com/labstack/echo/v4"
)

// init ...
func init() {
	fmt.Println("HELLO")

	config.Init()
	database.Connect()
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
