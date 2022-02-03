package view

import (
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func LogoutUser(c echo.Context) error {
	r := new(Response)

	if err := view.DeleteSession(c); err != nil {
		return c.JSON(http.StatusInternalServerError, r)
	}
	r.Message = "ログアウト成功"

	return c.JSON(http.StatusOK, r)
}
