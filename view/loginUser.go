package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	u := new(controller.ResponseUser)

	id, err := view.GetUserIDBySession(c)
	if err != nil {
		return c.JSON(http.StatusOK, u)
	}
	if u, err = controller.UserLoginSpecific(id); err != nil {
		return c.JSON(http.StatusInternalServerError, u)
	}

	return c.JSON(http.StatusOK, u)
}
