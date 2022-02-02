package controller

import "TsunoKento/emotionSNS/model"

func PostAll() (*[]model.PostWithUser, error) {
	p, err := model.GetAllPostWithUser()
	return p, err
}
