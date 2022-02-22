package main

import (
	"TsunoKento/emotionSNS/view"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("WEB_SERVER_URL")},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	env := os.Getenv("ENV")
	if env == "prod" {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(os.Getenv("HOST"))
		e.AutoTLSManager.Cache = autocert.DirCache("./certs")
		e.Pre(middleware.HTTPSWWWRedirect())
	}

	e.GET("/post/all", view.AllPost)
	e.GET("/post/all/:uid", view.AllPost)
	e.POST("/post/add", view.AddPost)
	e.GET("/user/login/google", view.GoogleLogin)
	e.GET("/user/login/google/callback", view.CallbackGoogleLogin)
	e.GET("/user/getUser/:uid", view.UserGet)
	e.POST("/user/profile/change", view.ChangeProfile)
	e.POST("/user/logout", view.LogoutUser)
	e.POST("/user/loginUser", view.LoginUser)
	e.POST("/like", view.ToggleLike)

	switch env {
	case "prod":
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	case "dev":
		e.Logger.Fatal(e.Start(":80"))
	default:
		e.Logger.Fatal(e.Start(":8000"))
	}
}
