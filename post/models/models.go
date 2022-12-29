package models

type Post struct {
	Author  string
	Title   string
	Content string
}

type RequestCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
