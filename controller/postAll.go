package controller

import (
	"TsunoKento/emotionSNS/controller/pkg"
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

	if len(*p) == 0 {
		return p, err
	}

	slice := *p
	for i := range slice {
		if slice[i].UserImage != "" {
			ui, _ := pkg.GetS3ImageEncode(slice[i].UserImage)
			slice[i].UserImage = ui
		}
	}

	return p, err
}
