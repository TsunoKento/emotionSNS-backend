package view

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//sessionからログインしているユーザーのIDを返却する
func GetUserIDBySession(c echo.Context) (uint, error) {
	sess, err := session.Get("session", c)
	if _, ok := sess.Values["id"]; !ok {
		return 0, errors.New("ログインしていません")
	}
	return sess.Values["id"].(uint), err
}

//sessionにユーザーのIDを設定する
func SetUserIDToSession(id uint, c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["id"] = id
	sess.Save(c.Request(), c.Response())
}

func DeleteSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}
