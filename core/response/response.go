package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c echo.Context, data interface{}) error {
	resp := new(response)
	resp.Status = http.StatusOK
	resp.Message = "Success"
	resp.Data = data

	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(http.StatusOK, resp)
}

func NotFound(c echo.Context) error {
	resp := new(response)
	resp.Status = http.StatusNotFound
	resp.Message = "Not Found"

	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(http.StatusNotFound, resp)
}

func InternalError(c echo.Context) error {
	resp := new(response)
	resp.Status = http.StatusInternalServerError
	resp.Message = "Internal Server Error"

	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(http.StatusInternalServerError, resp)
}

func ValidationError(c echo.Context, message interface{}) error {
	resp := new(response)
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = message

	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(http.StatusUnprocessableEntity, resp)
}
