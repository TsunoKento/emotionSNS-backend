package model

type Like struct {
	ID     uint
	UserID uint
	PostID uint
}

//いいねを追加します
func (l *Like) AddLike() (*Like, error) {
	r := db.Create(&l)
	return l, r.Error
}

//いいねを削除します
func (l *Like) DeleteLike() (*Like, error) {
	r := db.Where("user_id = ? AND post_id = ?", l.UserID, l.PostID).Delete(l)
	return l, r.Error
}
