package validations

import (
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Validator ...
type Validator struct{}

// ValidateObjectID ...
func (v Validator) ValidateObjectID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")

		// ValidateObjectID
		err := util.ValidateObjectID(id)
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("id", id)

		return next(c)
	}
}
