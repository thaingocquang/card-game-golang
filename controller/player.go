package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Player ...
type Player struct{}

// MyProfile godoc
// @Summary      get my profile
// @Description  player get my profile
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/profile [get]
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

// UpdatePlayer godoc
// @Summary      update player by id
// @Description  player update player by id
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        bot  body      dto.PlayerUpdate  true  "update bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /api/players [put]
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

// GetByID godoc
// @Summary      get player by id
// @Description  admin get player by id
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "player ID"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/profiles/{id} [get]
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

// GetListProfile godoc
// @Summary      get list player
// @Description  admin get list player
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/profiles [get]
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

// UpdateProfileByID godoc
// @Summary      update profile by id
// @Description  admin update profile by id
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "player ID"
// @Param        profile  body      dto.ProfileUpdate  true  "update bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/profiles/{id} [put]
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

// DeleteByID godoc
// @Summary      delete player by id
// @Description  admin delete player by id
// @Tags         players
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "player ID"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/profiles/{id} [delete]
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
