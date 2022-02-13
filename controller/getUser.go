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

	img, err := pkg.GetS3ImageEncode(u.Image)
	if err != nil {
		return ru, err
	}

	ru.UserID = u.UserID
	ru.Name = u.Name
	ru.Image = img

	return ru, nil
}
