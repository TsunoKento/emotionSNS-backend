package user

import (
	"TsunoKento/emotionSNS/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type ResponseUser struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Image  string `json:"image"`
}

func CurrentUser(c echo.Context) error {
	u := new(ResponseUser)

	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, u)
	}

	if a := sess.Values["auth"]; a == true {
		id := sess.Values["id"].(uint)
		user, _ := model.SearchByID(id)
		u.UserID = user.UserID
		u.Name = user.Name
		u.Image = user.Image
	}

	return c.JSON(http.StatusOK, u)
}
