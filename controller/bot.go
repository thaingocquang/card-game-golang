package controller

import (
	"card-game-golang/dto"
	"card-game-golang/util"
	"github.com/labstack/echo/v4"
)

// Bot ...
type Bot struct{}

// Create godoc
// @Summary      create bot
// @Description  create bot
// @Tags         bots
// @Accept       json
// @Produce      json
// @Param        account  body      dto.Bot  true  "create bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Router       /admin/bots [post]
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

// GetByID godoc
// @Summary      get bot by ID
// @Description  get bot by ID
// @Tags         bots
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "bot ID"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/bots/{id} [get]
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

// GetList godoc
// @Summary      get list bot
// @Description  get list bot by
// @Tags         bots
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        page  query      int     true  "page query"
// @Param        limit  query      int     true  "limit query"
// @Success      200  {object}  util.ResponsePaging
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/bots [get]
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

// UpdateByID godoc
// @Summary      update bot by id
// @Description  update bot by id
// @Tags         bots
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "bot ID"
// @Param        bot  body      dto.Bot  true  "update bot"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/bots/{id} [put]
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

// DeleteByID godoc
// @Summary      delete bot by id
// @Description  delete bot by id
// @Tags         bots
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param        id path string true "bot ID"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Security     ApiKeyAuth
// @Router       /admin/bots/{id} [delete]
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
