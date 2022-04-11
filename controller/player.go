package controller

import (
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Player ...
type Player struct{}

// MyProfile ...
func (p Player) MyProfile(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	profile, err := playerService.MyProfile(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, profile, "")
}

// Register ...
func (p Player) Register(c echo.Context) error {
	return nil
}

// Login ...
func (p Player) Login(c echo.Context) error {
	return nil
}

// UpdateMyProfile ...
func (p Player) UpdateMyProfile(c echo.Context) error {
	return nil
}

// UpdateMyPassword ...
func (p Player) UpdateMyPassword(c echo.Context) error {
	return nil
}
