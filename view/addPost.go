package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Content string `json:"content"`
	Emotion uint   `json:"emotion"`
}

func AddPost(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if req.Content == "" {
		return errors.New("投稿内容が空では登録できません")
	}

	id, err := view.GetUserIDBySession(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := controller.PostAdd(req.Content, req.Emotion, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
