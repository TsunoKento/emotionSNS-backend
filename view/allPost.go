package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AllPost(c echo.Context) error {
	id, err := view.GetUserIDBySession(c)
	if err != nil && err.Error() != "ログインしていません" {
		return c.String(http.StatusBadRequest, err.Error())
	}

	pd, err := controller.PostAll(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, pd)
}
