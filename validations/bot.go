package validations

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Bot ...
type Bot struct{}

// Create ...
func (b Bot) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// player register body
		var body dto.Bot

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.ValidateCreate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}

// Update ...
func (b Bot) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// player register body
		var body dto.Bot

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.ValidateUpdate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}
