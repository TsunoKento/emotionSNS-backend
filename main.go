package main

import (
	"TsunoKento/emotionSNS/auth"
	"TsunoKento/emotionSNS/post"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.GET("/api/postAll", post.ShowAll)
	e.POST("/api/addPost", post.AddPost)
	e.GET("/api/auth/google", auth.GoogleLogin)
	e.GET("/callback", auth.GoogleCallback)
	e.POST("/api/auth/logout", auth.Logout)

	e.Logger.Fatal(e.Start(":8000"))
}
