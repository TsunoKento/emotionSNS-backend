package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func Logout(c echo.Context) error {
	r := new(Response)

	sess, err := session.Get("session", c)
	if err != nil {
		r.Message = err.Error()
		return c.JSON(http.StatusForbidden, r)
	}
	fmt.Println(sess.Values["auth"])
	if a := sess.Values["auth"]; a == true {
		sess.Values["auth"] = false
		sess.Save(c.Request(), c.Response())
		r.Message = "ログアウト成功"
		return c.JSON(http.StatusOK, r)
	} else {
		fmt.Println("すでにログアウトしています")
	}
	return c.JSON(http.StatusOK, r)
}
