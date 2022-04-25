package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"fmt"
	"github.com/labstack/echo/v4"
)

// Game ...
type Game struct{}

// PlayByBotID ...
func (g Game) PlayByBotID(c echo.Context) error {
	var botID = c.Get("id").(string)
	var body = c.Get("body").(dto.GameVal)

	// jwtPayload for get my id
	jwtPayload, _ := util.GetJWTPayload(c)
	myID := jwtPayload["id"].(string)

	// process data
	gameJSON, err := gameService.PlayByBotID(body, botID, myID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, gameJSON, "")
}

// PlayRandom ...
func (g Game) PlayRandom(c echo.Context) error {
	var body = c.Get("body").(dto.GameVal)

	// jwtPayload for get my id
	jwtPayload, _ := util.GetJWTPayload(c)
	myID := jwtPayload["id"].(string)

	// process data
	gameBSON, err := gameService.PlayRandom(body, myID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, gameBSON, "")
}

// RecentGame ...
func (g Game) RecentGame(c echo.Context) error {
	return nil
}

// GetList ...
func (g Game) GetList(c echo.Context) error {
	paging := c.Get("paging").(util.Paging)

	// process
	games, err := gameService.GetList(&paging)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	fmt.Println(games)

	return util.Response200Paging(c, games, paging, "")
}
