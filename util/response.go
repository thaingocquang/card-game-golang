package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Response ...
type Response map[string]interface{}

// generateResponse ...
func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

// Response200 success ...
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "success"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

// Response400 bad request ...
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "bad request"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response401 unauthorized ...
func Response401(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "unauthorized"
	}
	return c.JSON(http.StatusUnauthorized, generateResponse(data, message))
}

// Response404 not found ...
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "not found"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}
