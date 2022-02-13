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

	uid := c.Param("uid")

	pd, err := controller.PostAll(id, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if len(*pd) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, pd)
}
