package validations

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Player ...
type Player struct{}

// Register ...
func (p Player) Register(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// player register body
		var body dto.Player

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}

// Login ...
func (p Player) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// player login body
		var body dto.PlayerLogin

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}
