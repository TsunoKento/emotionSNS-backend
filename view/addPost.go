package view

import (
	"TsunoKento/emotionSNS/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Content  string `json:"content"`
	Emotions []uint `json:"emotions"`
}

func AddPost(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return err
	}
	res, err := controller.PostAdd(req.Content, req.Emotions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
