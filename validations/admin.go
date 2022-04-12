package validations

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Admin ...
type Admin struct{}

// Login ...
func (a Admin) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// admin login body
		var body dto.Admin

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

// CheckRole ...
func (a Admin) CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtPayload, _ := util.GetJWTPayload(c)

		if jwtPayload["admin"] == true {
			return next(c)
		}

		return util.Response400(c, nil, "authorization fail: not admin")
	}
}
