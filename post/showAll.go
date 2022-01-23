package post

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Imame string `json:"image,omitempty"`
}

type Post struct {
	ID      string `json:"id"`
	Image   string `json:"image,omitempty"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Like    int64  `json:"like"`
}

type PostData struct {
	Type int64    `json:"type"`
	User User     `json:"user"`
	Post Post     `json:"post"`
	Tags []string `json:"tags,omitempty"`
}

func ShowAll(c echo.Context) error {
	pd := []PostData{{
		Type: 0,
		User: User{
			ID:    "asdf123",
			Name:  "kento",
			Imame: "https://source.unsplash.com/bIhpiQA009k",
		},
		Post: Post{
			ID:      "1",
			Image:   "https://source.unsplash.com/brFsZ7qszSY",
			Date:    "2021年12月29日",
			Content: "今日も犬は可愛い",
			Like:    9,
		},
		Tags: []string{"#今日のワンコ", "#ペット"},
	}, {
		Type: 0,
		User: User{
			ID:   "tarousama",
			Name: "太郎",
		},
		Post: Post{
			ID:      "2",
			Date:    "2021年11月21日",
			Content: "肌綺麗だって褒められた嬉しい",
			Like:    4,
		},
	}, {
		Type: 1,
		User: User{
			ID:    "asdf123",
			Name:  "kento",
			Imame: "https://source.unsplash.com/bIhpiQA009k",
		},
		Post: Post{
			ID:      "3",
			Date:    "2021年12月21日",
			Content: "チャンネル争いに負けた",
			Like:    0,
		},
	}, {
		Type: 2,
		User: User{
			ID:    "asdf123",
			Name:  "kento",
			Imame: "https://source.unsplash.com/bIhpiQA009k",
		},
		Post: Post{
			ID:      "4",
			Date:    "2021年12月30日",
			Content: "お米炊くの忘れてた....",
			Like:    4,
		}},
		{
			Type: 3,
			User: User{
				ID:    "asdf123",
				Name:  "kento",
				Imame: "https://source.unsplash.com/bIhpiQA009k",
			},
			Post: Post{
				ID:      "5",
				Date:    "2022年1月1日",
				Content: "新年早々面白すぎwww",
				Like:    21,
			}},
		{
			Type: 3,
			User: User{
				ID:   "tarousama",
				Name: "太郎",
			},
			Post: Post{
				ID:      "6",
				Date:    "2020年10月19日",
				Content: "今週のジャンプ面白い",
				Like:    5,
			},
			Tags: []string{"#ジャンプ"},
		},
		{
			Type: 4,
			User: User{
				ID:    "asdf123",
				Name:  "kento",
				Imame: "https://source.unsplash.com/bIhpiQA009k",
			},
			Post: Post{
				ID:      "7",
				Image:   "https://source.unsplash.com/TobZaa8ZwI4",
				Date:    "2021年11月13日",
				Content: "暇",
				Like:    0,
			},
		},
	}
	return c.JSON(http.StatusOK, pd)
}
