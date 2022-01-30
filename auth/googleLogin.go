package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	provider, _ = oidc.NewProvider(context.Background(), "https://accounts.google.com")
	conf        = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8000/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		Endpoint:     provider.Endpoint(),
	}
	verifier    = provider.Verifier(&oidc.Config{ClientID: os.Getenv("GOOGLE_CLIENT_ID")})
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

type User struct {
	ID           uint
	ThirdPartyID string
	UserID       string
	Name         string
	Image        string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GoogleLogin(c echo.Context) error {
	randomState = randString(8)
	url := conf.AuthCodeURL(randomState)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
	if c.FormValue("state") != randomState {
		fmt.Println("有効なstateがありません")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	token, err := conf.Exchange(context.Background(), c.FormValue("code"))
	if err != nil {
		fmt.Printf("トークンを取得できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		fmt.Println("IDトークンを取得できませんでした")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	idToken, err := verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		fmt.Printf("IDトークンを検証できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	var claims struct {
		Subject string `json:"sub"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}
	if err := idToken.Claims(&claims); err != nil {
		fmt.Printf("JSONを検証できませんでした: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}
	fmt.Println(claims)

	dsn := "root:root@tcp(db:3306)/emotion_sns?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("DBアクセスに失敗しました: %s\n", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}

	var user User
	result := db.Where("third_party_id = ?", claims.Subject).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		uid := randString(8)
		user = User{ThirdPartyID: claims.Subject, UserID: uid, Name: claims.Name, Email: claims.Email, Image: claims.Picture}
		db.Create(&user)
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["id"] = user.ID
	sess.Values["auth"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
