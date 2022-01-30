package controller

type Response struct {
	Content  string `json:"content"`
	Emotions []uint `json:"emotions"`
	Status   string `json:"status"`
}

func PostAdd(content string, emotions []uint) (*Response, error) {
	r := new(Response)
	r.Content = content
	r.Emotions = emotions
	r.Status = "success"

	return r, nil
}
