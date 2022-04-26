package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Auth ...
type Auth struct{}

// PlayerRegister godoc
// @Summary      register an account
// @Description  registering an account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        account  body      dto.Player  true  "register account"
// @Success      200  {object}  util.ResponseTest
// @Failure      400  {object}  util.Response
// @Router       /api/register [post]
func (a Auth) PlayerRegister(c echo.Context) error {
	var player = c.Get("body").(dto.Player)

	// process data
	if err := authService.Register(player); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, nil, "")
}

// PlayerLogin godoc
// @Summary      player login
// @Description  player login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        account  body      dto.PlayerLogin  true  "login account"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Router       /api/login [post]
func (a Auth) PlayerLogin(c echo.Context) error {
	var player = c.Get("body").(dto.PlayerLogin)

	// process data
	token, err := authService.Login(player)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token": token,
	}

	// success
	return util.Response200(c, data, "")
}

// AdminLogin godoc
// @Summary      admin login
// @Description  admin login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        account  body      dto.Admin  true  "login account"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Router       /admin/login [post]
func (a Auth) AdminLogin(c echo.Context) error {
	var admin = c.Get("body").(dto.Admin)

	// process data
	token, err := authService.AdminLogin(admin)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token": token,
	}

	return util.Response200(c, data, "")
}
