package validations

import (
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Validator ...
type Validator struct{}

// ValidateObjectID ...
func (Validator) ValidateObjectID(next echo.HandlerFunc) echo.HandlerFunc {
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

// Paging ...
func (Validator) Paging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var paging util.Paging

		if err := c.Bind(&paging); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		if paging.Page < 0 && paging.Limit < 0 {
			return util.Response400(c, nil, "page or limit must great than or equal 0")
		}

		// validate Paging attribute (page, limit)

		c.Set("paging", paging)

		return next(c)
	}
}
