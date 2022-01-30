package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GoogleLogin(c echo.Context) error {
	url, _ := controller.SetLoginUrl()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackGoogleLogin(c echo.Context) error {
	state := c.FormValue("state")
	code := c.FormValue("code")

	user, err := controller.CallbackGoogleLogin(state, code)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "http://localhost:3000")
	}

	view.SetUserIDToSession(user.ID, c)

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
