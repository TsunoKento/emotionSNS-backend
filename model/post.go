package model

import (
	"time"
)

type Post struct {
	ID          uint
	UserID      uint
	Content     string
	Image       string
	PublishedAt time.Time `gorm:"autoCreateTime"`
	EmotionID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//新規投稿をデータベースに格納する
func (p *Post) CreatePost(uid, eid uint, con string) error {
	p = &Post{UserID: uid, EmotionID: eid, Content: con}
	r := db.Create(&p)
	return r.Error
}
