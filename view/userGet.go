package view

import (
	"TsunoKento/emotionSNS/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserGet(c echo.Context) error {
	uid := c.Param("uid")

	u, err := controller.GetUser(uid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ユーザー情報を取得できませんでした")
	}

	return c.JSON(http.StatusOK, u)
}
