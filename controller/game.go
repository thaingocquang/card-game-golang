package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
	"strconv"
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
	gameBSON, err := gameService.PlayByBotID(body, botID, myID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, gameBSON, "")
}

// PlayRandom ...
func (g Game) PlayRandom(c echo.Context) error {
	return nil
}

// RecentGame ...
func (g Game) RecentGame(c echo.Context) error {
	return nil
}

// GetList ...
func (g Game) GetList(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// process
	games, totalDocs, err := gameService.GetList(page, limit)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"list": games,
		"paginationInfo": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": totalDocs,
		},
	}

	return util.Response200(c, data, "")
}
