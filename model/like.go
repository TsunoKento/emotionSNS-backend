package model

type Like struct {
	ID     uint
	UserID uint
	PostID uint
}

//いいねを追加します
func AddLike(l *Like) (*Like, error) {
	r := db.Create(&l)
	return l, r.Error
}

//いいねを削除します
func DeleteLike(l *Like) (*Like, error) {
	r := db.Where("user_id = ? AND post_id = ?", l.UserID, l.PostID).Delete(l)
	return l, r.Error
}
