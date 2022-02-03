package controller

import (
	"TsunoKento/emotionSNS/model"
)

func PostAll() (*[]model.PostWithUserWithLikes, error) {
	p, err := model.GetAllPostWithUser()

	return p, err
}
