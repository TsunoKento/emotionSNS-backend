package controller

import (
	"TsunoKento/emotionSNS/model"
)

func PostAll(id uint, uid string) (*model.SlicePostWithUserWithLikes, error) {
	p := new(model.SlicePostWithUserWithLikes)
	var err error
	if uid == "" {
		err = p.GetAllPostWithUser(id)
	} else {
		err = p.GetAllPostWithUserWhereUserID(id, uid)
	}

	return p, err
}
