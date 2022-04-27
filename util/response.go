package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//Response ...
//type Response map[string]interface{}

// Response ...
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ResponsePaging ...
type ResponsePaging struct {
	Data    interface{} `json:"data"`
	Paging  interface{} `json:"paginationInfo"`
	Message string      `json:"message"`
}

// generateResponseTest ...
func generateResponse(data interface{}, message string) Response {
	return Response{
		Data:    data,
		Message: message,
	}
}

//// generateResponse ...
//func generateResponse(data interface{}, message string) Response {
//	return Response{
//		"data":    data,
//		"message": message,
//	}
//}

func generateResponsePaging(data, paging interface{}, message string) ResponsePaging {
	return ResponsePaging{
		Data:    data,
		Paging:  paging,
		Message: message,
	}
}

// Response200 success ...
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "success"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

// Response200Paging success with paging ...
func Response200Paging(c echo.Context, data, paging interface{}, message string) error {
	if message == "" {
		message = "success"
	}
	return c.JSON(http.StatusOK, generateResponsePaging(data, paging, message))
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

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
