package model

import "time"

type PostWithUserWithLikes struct {
	PostID      uint      `json:"postId"`
	Content     string    `json:"content"`
	PostImage   string    `json:"postImage"`
	PublishedAt time.Time `json:"publishedAt"`
	EmotionID   uint      `json:"emotionId"`
	UserID      string    `json:"userId"`
	Name        string    `json:"name"`
	UserImage   string    `json:"userImage"`
	LikeFlag    bool      `json:"likeFlag"` //0ならfalse, 1ならtrueになる
	LikeCount   uint      `json:"likeCount"`
}

type SlicePostWithUserWithLikes []PostWithUserWithLikes

//すべての投稿を取得してログイン中のユーザーがいいねを押しているかも返す
func (p *SlicePostWithUserWithLikes) GetAllPostWithUser(uid uint) error {
	//p := new([]PostWithUserWithLikes)
	r := db.Table("posts").
		Select("posts.id AS post_id, posts.content, posts.image AS post_image, posts.published_at, posts.emotion_id, users.user_id, users.name, users.image AS user_image, COALESCE(flag, ?) AS like_flag, COALESCE(count, ?) AS like_count", 0, 0).
		Joins("INNER JOIN users ON posts.user_id = users.id").
		Joins("LEFT OUTER JOIN (?) AS f ON posts.id = f.post_id", db.Table("likes").Select("post_id, 1 AS flag").Where("user_id = ?", uid)).
		Joins("LEFT OUTER JOIN (?) AS l ON posts.id = l.post_id", db.Table("likes").Select("post_id, COUNT(*) AS count").Group("post_id")).
		Order("published_at DESC").
		Scan(&p)
	return r.Error
}
