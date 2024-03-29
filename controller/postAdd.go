package controller

import "TsunoKento/emotionSNS/model"

type Response struct {
	Content string `json:"content"`
	Emotion uint   `json:"emotions"`
	Status  string `json:"status"`
}

func PostAdd(content string, emotion, id uint) (*Response, error) {
	r := new(Response)

	post := new(model.Post)
	err := post.CreatePost(id, emotion, content)
	if err != nil {
		return r, err
	}

	r.Status = "success"
	return r, nil
}
