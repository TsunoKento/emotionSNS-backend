package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ChangeProfile(c echo.Context) error {
	var req struct {
		Image  string `json:"image"`
		UserID string `json:"userId"`
		Name   string `json:"name"`
	}
	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, "リクエストの型にあいません")
	}

	id, err := view.GetUserIDBySession(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = controller.ProfileChange(req.Image, req.UserID, req.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if req.UserID != "" {
		return c.JSON(http.StatusSeeOther, os.Getenv("WEB_SERVER_URL")+"/profile/"+req.UserID)
	}
	return c.NoContent(http.StatusOK)
}
