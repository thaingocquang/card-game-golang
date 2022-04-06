package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Auth ...
type Auth struct{}

// PlayerRegister ...
func (a Auth) PlayerRegister(c echo.Context) error {
	var player = c.Get("body").(dto.Player)

	// process data
	if err := playerService.Register(player); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

// PlayerLogin ...
func (a Auth) PlayerLogin(c echo.Context) error {
	return nil
}

// AdminLogin ...
func (a Auth) AdminLogin(c echo.Context) error {
	return nil
}
