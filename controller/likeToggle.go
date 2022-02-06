package controller

import "TsunoKento/emotionSNS/model"

func LikeToggle(f bool, uid, pid uint) error {
	like := new(model.Like)
	like.UserID = uid
	like.PostID = pid
	if f {
		if _, err := like.AddLike(); err != nil {
			return err
		}
	} else {
		if _, err := like.DeleteLike(); err != nil {
			return err
		}
	}
	return nil
}
