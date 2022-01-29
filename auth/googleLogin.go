package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8000/callback",
		Scopes:       []string{"profile"},
		Endpoint:     google.Endpoint,
	}
	randomState = randString(8)
)

func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type GoogleUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func GoogleLogin(c echo.Context) error {
	url := conf.AuthCodeURL(randomState)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
	if c.FormValue("state") != randomState {
		fmt.Println("有効なstateがありません")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	randomState = randString(8)
	token, err := conf.Exchange(context.Background(), c.FormValue("code"))
	if err != nil {
		fmt.Printf("トークンを取得できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("GETリクエストを作成できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("レスポンスを解析できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	gu := new(GoogleUser)
	err = json.Unmarshal(content, &gu)
	if err != nil {
		fmt.Printf("取得データを変換できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	//TODO GoogleIDの有無で過去にアカウントを作ったことがあるか判断、無ければデータベースに登録する
	fmt.Println(gu.ID)
	fmt.Println(gu.Name)
	fmt.Println(gu.Picture)

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userId"] = "kentots"
	sess.Values["auth"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
