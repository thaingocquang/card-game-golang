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
	profile, err := playerService.GetByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, profile, "")
}

// UpdatePlayer ...
func (p Player) UpdatePlayer(c echo.Context) error {
	// player id & update body
	var id = c.Get("id").(string)
	var body = c.Get("body").(dto.PlayerUpdate)

	// process
	err := playerService.Update(id, body)

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

func (p Player) GetByID(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	profile, err := playerService.GetByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, profile, "")
}

// GetListProfile ...
func (p Player) GetListProfile(c echo.Context) error {
	paging := c.Get("paging").(util.Paging)

	// fulfill
	paging.Fulfill()

	// process
	profiles, err := playerService.GetListProfile(&paging)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200Paging(c, profiles, paging, "")
}

// UpdateProfileByID ...
func (p Player) UpdateProfileByID(c echo.Context) error {
	// player id & update body
	var id = c.Get("id").(string)
	var body = c.Get("body").(dto.ProfileUpdate)

	// process
	err := playerService.UpdateProfile(id, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, id, "")
}

// DeleteByID ...
func (p Player) DeleteByID(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	err := playerService.DeleteByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

// DeleteAll ...
func (p Player) DeleteAll(c echo.Context) error {
	// process
	err := playerService.DeleteAll()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

//// Test ...
//func (p Player) Test(c echo.Context) error {
//	// process
//	err := playerService.Test()
//	if err != nil {
//		return util.Response400(c, nil, err.Error())
//	}
//
//	// success
//	return util.Response200(c, nil, "")
//}
