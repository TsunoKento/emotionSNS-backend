package controller

import (
	"TsunoKento/emotionSNS/controller/pkg"
	"TsunoKento/emotionSNS/model"
)

func GetUser(uid string) (*ResponseUser, error) {
	ru := new(ResponseUser)
	u := new(model.User)

	if err := u.SearchByUserID(uid); err != nil {
		return ru, err
	}

	if u.Image != "" {
		img, err := pkg.GetS3ImageEncode(u.Image)
		if err != nil {
			return ru, err
		}
		ru.Image = img
	}

	ru.UserID = u.UserID
	ru.Name = u.Name

	return ru, nil
}
