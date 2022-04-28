package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Game ...
type Game struct{}

// PlayByBotID godoc
// @Summary      play game by id
// @Description  player play game by id
// @Tags         games
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "bot ID"
// @Param        bot  body      dto.GameVal  true  "update bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/games/{id} [post]
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

// PlayRandom godoc
// @Summary      play game random
// @Description  player play game random
// @Tags         games
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        bot  body      dto.GameVal  true  "update bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/games [post]
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

// RecentGame godoc
// @Summary      get recent game
// @Description  player get recent game
// @Tags         games
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/games [get]
func (g Game) RecentGame(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	recentGames, err := gameService.Recent(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, recentGames, "")
}

// GetList godoc
// @Summary      get list game
// @Description  admin get list game
// @Tags         games
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Success      200  {object}  util.ResponsePaging
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/games [get]
func (g Game) GetList(c echo.Context) error {
	paging := c.Get("paging").(util.Paging)

	// fulfill
	paging.Fulfill()

	// process
	games, err := gameService.GetList(&paging)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200Paging(c, games, paging, "")
}
