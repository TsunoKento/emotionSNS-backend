package controller

import (
	"TsunoKento/emotionSNS/model"
	"TsunoKento/emotionSNS/pkg"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var (
	provider, _ = oidc.NewProvider(context.Background(), "https://accounts.google.com")
	conf        = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8000/user/login/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		Endpoint:     provider.Endpoint(),
	}
	verifier    = provider.Verifier(&oidc.Config{ClientID: os.Getenv("GOOGLE_CLIENT_ID")})
	randomState = pkg.RandString(8)
)

func SetLoginUrl() (string, error) {
	randomState = pkg.RandString(8)
	url := conf.AuthCodeURL(randomState)
	return url, nil
}

func CallbackGoogleLogin(state, code string) (*model.User, error) {
	if state != randomState {
		return nil, errors.New("有効なstateがありません")
	}

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		fmt.Println("IDトークンを取得できませんでした")
		return nil, errors.New("IDトークンを取得できませんでした")
	}

	idToken, err := verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		return nil, err
	}

	var claims struct {
		Subject string `json:"sub"`
		Name    string `json:"name"`
		Email   string `json:"email"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	user := new(model.User)
	err = user.SearchByThirdPartyID(claims.Subject)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		uid := pkg.RandString(12)
		if err := user.CreateUser(claims.Subject, uid, claims.Name, claims.Email); err != nil {
			return user, err
		}
	}

	return user, nil
}
