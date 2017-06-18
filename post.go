package main

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPost(title, content string) *Post {
	return &Post{
		Title:   title,
		Content: content,
	}
}
