package dto

import (
	"libra.com/blog/internal/data/blogs"
)

type BlogInfo struct {
	Id       int64       `json:"id"`
	Title    string      `json:"title"`
	Desc     string      `json:"desc"`
	Author   string      `json:"author"`
	Tags     []blogs.Tag `json:"tags"`
	Category []blogs.Tag `json:"category"`
	CreateAt int64       `json:"createAt"`
	UpdateAt int64       `json:"updateAt"`
	ImgUrl   string      `json:"imgUrl"`
}

type BlogList struct {
	Total int64       `json:"total"`
	Blogs []*BlogInfo `json:"list"`
}
