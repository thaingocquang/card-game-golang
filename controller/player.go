package controller

import (
	"card-game-golang/dto"
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

// UpdateMyProfile ...
func (p Player) UpdateMyProfile(c echo.Context) error {
	// player id & update body
	var id = c.Get("id").(string)
	var body = c.Get("body").(dto.PlayerUpdate)

	// process
	err := playerService.UpdateProfile(id, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, id, "")
}

// UpdateMyPassword ...
func (p Player) UpdateMyPassword(c echo.Context) error {
	return nil
}
