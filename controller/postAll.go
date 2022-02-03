package controller

import (
	"TsunoKento/emotionSNS/model"
)

func PostAll(id uint) (*[]model.PostWithUserWithLikes, error) {
	p, err := model.GetAllPostWithUser(id)

	return p, err
}
