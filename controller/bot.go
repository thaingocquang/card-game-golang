package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Bot ...
type Bot struct{}

// Create ...
func (b Bot) Create(c echo.Context) error {
	var body = c.Get("body").(dto.Bot)

	// process
	err := botService.Create(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// GetByID ...
func (b Bot) GetByID(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	bot, err := botService.GetByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bot, "")
}

//// GetList ...
//func (b Bot) GetList(c echo.Context) error {
//	page, _ := strconv.Atoi(c.QueryParam("page"))
//	limit, _ := strconv.Atoi(c.QueryParam("limit"))
//
//	// process
//	bots, totalDocs, err := botService.GetList(page, limit)
//	if err != nil {
//		return util.Response400(c, nil, err.Error())
//	}
//
//	data := map[string]interface{}{
//		"list": bots,
//		"paginationInfo": map[string]interface{}{
//			"page":  page,
//			"limit": limit,
//			"total": totalDocs,
//		},
//	}
//
//	return util.Response200(c, data, "")
//}

// GetList ...
func (b Bot) GetList(c echo.Context) error {
	paging := c.Get("paging").(util.Paging)

	// fulfill
	paging.Fulfill()

	// process
	bots, err := botService.GetList(&paging)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200Paging(c, bots, paging, "")
}

// UpdateByID ...
func (b Bot) UpdateByID(c echo.Context) error {
	var id = c.Get("id").(string)
	var body = c.Get("body").(dto.Bot)

	// process
	err := botService.UpdateByID(id, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// DeleteByID ...
func (b Bot) DeleteByID(c echo.Context) error {
	var id = c.Get("id").(string)

	// process
	err := botService.DeleteByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// DeleteAll ...
func (b Bot) DeleteAll(c echo.Context) error {
	// process
	err := botService.DeleteAll()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}
