package controller

import (
	"TsunoKento/emotionSNS/model"
)

func PostAll(id uint) (*model.SlicePostWithUserWithLikes, error) {
	p := new(model.SlicePostWithUserWithLikes)
	err := p.GetAllPostWithUser(id)

	return p, err
}
