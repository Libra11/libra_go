package dto

import (
	"libra.com/blog/internal/data/blogs"
)

type BlogInfo struct {
	Id       int64            `json:"id"`
	Title    string           `json:"title"`
	Desc     string           `json:"desc"`
	Author   string           `json:"author"`
	Tags     []blogs.Tag      `json:"tags"`
	Category []blogs.Category `json:"category"`
	CreateAt int64            `json:"createAt"`
	UpdateAt int64            `json:"updateAt"`
	ImgUrl   string           `json:"imgUrl"`
}

type BlogList struct {
	Total int64       `json:"total"`
	Blogs []*BlogInfo `json:"list"`
}

type WordInfo struct {
	Id         int64  `json:"id"`
	Word       string `json:"word"`
	CreateAt   int64  `json:"createAt"`
	UpdateAt   int64  `json:"updateAt"`
	Definition string `json:"definition"`
	Example    string `json:"example"`
	Phrase     string `json:"phrase"`
	Phonetic   string `json:"phonetic"`
}

type WordList struct {
	Total int64       `json:"total"`
	Words []*WordInfo `json:"list"`
}
