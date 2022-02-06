package controller

import "TsunoKento/emotionSNS/model"

type ResponseUser struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	Image  string `json:"image"`
}

func UserLoginSpecific(id uint) (*ResponseUser, error) {
	ru := new(ResponseUser)
	u := new(model.User)
	err := u.SearchByID(id)
	if err != nil {
		return ru, err
	}
	ru.UserID = u.UserID
	ru.Name = u.Name
	ru.Image = u.Image

	return ru, nil
}
