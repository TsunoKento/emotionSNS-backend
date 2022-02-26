package view

import (
	"TsunoKento/emotionSNS/controller"
	view "TsunoKento/emotionSNS/view/pkg"
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ChangeProfile(c echo.Context) error {
	var req struct {
		Image           string `json:"image"`
		UserID          string `json:"userId"`
		Name            string `json:"name"`
		IsUserIDChanged bool   `json:"isUserIdChanged"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "リクエストの型にあいません")
	}

	if req.UserID == "" || req.Name == "" {
		return errors.New("UserIDとNameは空白では登録できません")
	}

	if req.IsUserIDChanged && controller.CheckUserID(req.UserID) {
		return echo.NewHTTPError(http.StatusInternalServerError, "そのIDは既に使われています")
	}

	id, err := view.GetUserIDBySession(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = controller.ProfileChange(req.Image, req.UserID, req.Name, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var res struct {
		Url     string `json:"url"`
		Message string `json:"message"`
	}

	if req.IsUserIDChanged {
		res.Url = os.Getenv("WEB_SERVER_URL") + "/profile/" + req.UserID
		res.Message = "新しいページに遷移します"
		return c.JSON(http.StatusSeeOther, res)
	}
	res.Message = "変更しました"
	return c.JSON(http.StatusOK, res)
}
