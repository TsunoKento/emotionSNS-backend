package controller

import (
	"TsunoKento/emotionSNS/model"
	"errors"

	"gorm.io/gorm"
)

//対象のuseridがあればtrueなければfalseを返す
func CheckUserID(uid string) bool {
	u := new(model.User)

	if err := u.SearchByUserID(uid); errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
