package controller

import "TsunoKento/emotionSNS/model"

func GetUser(uid string) (*ResponseUser, error) {
	ru := new(ResponseUser)
	u := new(model.User)

	if err := u.SearchByUserID(uid); err != nil {
		return ru, err
	}

	ru.UserID = u.UserID
	ru.Name = u.Name
	ru.Image = u.Image

	return ru, nil
}
