package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ToggleLike(c echo.Context) error {
	var req struct {
		Like   bool `json:"like"`
		PostID uint `json:"postId"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの型に合わないです")
	}

	id, err := view.GetUserIDBySession(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := controller.LikeToggle(req.Like, id, req.PostID); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "いいね処理に成功しました")
}
