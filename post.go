package main

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPost(title, content string) *Post {
	return &Post{
		Title:   title,
		Content: content,
	}
}
