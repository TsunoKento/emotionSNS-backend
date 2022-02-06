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
func (u *User) CreateUser(tpid, uid, name, email, image string) error {
	u = &User{ThirdPartyID: tpid, UserID: uid, Name: name, Email: email, Image: image}
	result := db.Create(&u)
	return result.Error
}

//third_party_idから該当するユーザーを検索する
func (u *User) SearchByThirdPartyID(tpid string) error {
	result := db.Where("third_party_id = ?", tpid).Take(&u)
	return result.Error
}

//IDから該当するユーザーを検索する
func (u *User) SearchByID(id uint) error {
	result := db.First(&u, id)
	return result.Error
}
