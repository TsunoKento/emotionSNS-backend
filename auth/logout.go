package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

type Response struct {
	Message string `json:"message"`
}

func Logout(c echo.Context) error {
	r := new(Response)
	cookie, err := c.Cookie("session")
	if err != nil {
		r.Message = err.Error()
		return c.JSON(http.StatusBadRequest, r)
	}
	cookie.Path = "/"
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	r.Message = "ログアウト成功"
	return c.JSON(http.StatusOK, r)
}
