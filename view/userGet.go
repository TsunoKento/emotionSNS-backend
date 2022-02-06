package view

import (
	"TsunoKento/emotionSNS/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserGet(c echo.Context) error {
	var req struct {
		UserID string `json:"userId"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの型に合わないです")
	}

	u, err := controller.GetUser(req.UserID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ユーザー情報を取得できませんでした")
	}

	return c.JSON(http.StatusOK, u)
}
