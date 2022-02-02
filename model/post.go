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

type PostWithUser struct {
	PostID      uint      `json:"postId"`
	Content     string    `json:"content"`
	PostImage   string    `json:"postImage"`
	PublishedAt time.Time `json:"publishedAt"`
	EmotionID   uint      `json:"emotionId"`
	UserID      string    `json:"userId"`
	Name        string    `json:"name"`
	UserImage   string    `json:"userImage"`
}

//新規投稿をデータベースに格納する
func CreatePost(p *Post) (*Post, error) {
	r := db.Create(&p)
	return p, r.Error
}

//すべての投稿を取得する
func GetAllPostWithUser() (*[]PostWithUser, error) {
	pwu := new([]PostWithUser)
	r := db.Table("posts").
		Select("posts.id as post_id, posts.content, posts.image as post_image, posts.published_at, posts.emotion_id, users.user_id, users.name, users.image as user_image").
		Joins("inner join users on posts.user_id = users.id").
		Order("published_at desc").
		Scan(&pwu)
	return pwu, r.Error
}
