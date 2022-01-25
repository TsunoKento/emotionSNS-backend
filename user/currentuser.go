package user

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserID  string `json:"userId"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func CurrentUser(c echo.Context) error {
	u := new(User)

	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusOK, u)
	}

	if a := sess.Values["auth"]; a == false {
		return c.JSON(http.StatusOK, c)
	} else {
		u.UserID = "kentots"
		u.Name = "Kento"
		u.Picture = "https://images.unsplash.com/photo-1416339306562-f3d12fefd36f"
	}

	return c.JSON(http.StatusOK, u)
}
