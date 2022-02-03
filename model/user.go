package model

import (
	model "TsunoKento/emotionSNS/model/pkg"
	"time"
)

type User struct {
	ID           uint
	ThirdPartyID string
	UserID       string
	Name         string
	Image        string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var (
	db = model.Connect()
)

//ユーザー情報から新しくユーザーを登録する
func CreateUser(tpid, uid, name, email, image string) (*User, error) {
	user := User{ThirdPartyID: tpid, UserID: uid, Name: name, Email: email, Image: image}
	result := db.Create(&user)
	return &user, result.Error
}

//third_party_idから該当するユーザーを検索する
func SearchByThirdPartyID(tpid string) (*User, error) {
	user := new(User)
	result := db.Where("third_party_id = ?", tpid).Take(&user)
	return user, result.Error
}

//IDから該当するユーザーを検索する
func SearchByID(id uint) (*User, error) {
	user := new(User)
	result := db.First(&user, id)
	return user, result.Error
}
