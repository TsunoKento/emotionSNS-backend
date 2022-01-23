package post

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Content  string  `json:"content"`
	Emotions []int64 `json:"emotions"`
}

type Response struct {
	Content  string  `json:"content"`
	Emotions []int64 `json:"emotions"`
	Status   string  `json:"status"`
}

func AddPost(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}
	r := new(Response)
	r.Content = m.Content
	r.Emotions = m.Emotions
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}
