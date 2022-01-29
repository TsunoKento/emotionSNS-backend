package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Image  string `json:"image"`
}

func CurrentUser(c echo.Context) error {
	u := new(User)

	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK, u)
	}
	if a := sess.Values["auth"]; a == false {
		fmt.Println("ログインしていません")
		return c.JSON(http.StatusOK, c)
	} else {
		fmt.Println("ログイン情報を返却します")
		u.UserID = "kentots"
		u.Name = "Kento"
		u.Image = "https://images.unsplash.com/photo-1416339306562-f3d12fefd36f"
	}

	return c.JSON(http.StatusOK, u)
}
