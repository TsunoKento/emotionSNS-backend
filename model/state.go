package model

type State struct {
	ID    uint
	State string
}

//stateを登録する
func (s *State) CreateState(ns string) error {
	s = &State{State: ns}
	r := db.Create(&s)
	return r.Error
}

//stateを検索する
func (s *State) SearchState(ss string) error {
	r := db.Where("state = ?", ss).Take(&s)
	return r.Error
}

//stateを削除する
func (s *State) DeleteState() error {
	r := db.Delete(&s)
	return r.Error
}
