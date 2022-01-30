package controller

import "TsunoKento/emotionSNS/model"

type ResponseUser struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Image  string `json:"image"`
}

func UserLoginSpecific(id uint) (*ResponseUser, error) {
	u := new(ResponseUser)
	user, err := model.SearchByID(id)
	if err != nil {
		return u, err
	}
	u.UserID = user.UserID
	u.Name = user.Name
	u.Image = user.Image

	return u, nil
}
