package model

import "time"

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

func CreatePost(post *Post) (*Post, error) {
	result := db.Create(&post)
	return post, result.Error
}
