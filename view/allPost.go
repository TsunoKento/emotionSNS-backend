package view

import (
	"TsunoKento/emotionSNS/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AllPost(c echo.Context) error {
	pd, err := controller.PostAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, pd)
}
