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
	r := db.Create(&u)
	return r.Error
}

//third_party_idから該当するユーザーを検索する
func (u *User) SearchByThirdPartyID(tpid string) error {
	r := db.Where("third_party_id = ?", tpid).Take(&u)
	return r.Error
}

//IDから該当するユーザーを検索する
func (u *User) SearchByID(id uint) error {
	r := db.First(&u, id)
	return r.Error
}

//UserIDから該当するユーザーを検索する
func (u *User) SearchByUserID(uid string) error {
	r := db.Where("user_id = ?", uid).First(&u)
	return r.Error
}

//IDのユーザー情報を変更する
func (u *User) UpdateUser() error {
	r := db.Updates(u)
	return r.Error
}
