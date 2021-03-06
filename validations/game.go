package validations

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Game ...
type Game struct{}

// GameValue ...
func (g Game) GameValue(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// game value
		var body dto.GameVal

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
