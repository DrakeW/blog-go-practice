package main

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"path"
)

type PostsController struct {
	AppController
	*render.Render
	*sql.DB
}

func (c *PostsController) Index(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	rows, err := c.Query("SELECT title, content FROM posts")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer rows.Close()
	posts := make([]*Post, 0)
	for rows.Next() {
		var title, content string
		if err := rows.Scan(&title, &content); err != nil {
			log.Println(err.Error())
			return err
		}
		posts = append(posts, NewPost(title, content))
	}
	c.HTML(rw, http.StatusOK, "index", posts)
	return nil
}

func (c *PostsController) Show(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	var title, content string
	postId := p.ByName("id")
	err := c.QueryRow("SELECT title, content FROM posts WHERE rowid = ?", postId).Scan(&title, &content)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	post := NewPost(title, content)
	c.HTML(rw, http.StatusOK, "show", post)
	return nil
}

func (c *PostsController) New(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	post := NewPost("", "")
	err := c.HTML(rw, http.StatusOK, "new", post)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (c *PostsController) Create(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	title, content := r.FormValue("title"), r.FormValue("content")
	res, err := c.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	postId, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	http.Redirect(rw, r, fmt.Sprintf("/posts/%d", postId), http.StatusSeeOther)
	return nil
}

func (c *PostsController) Edit(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	postId := p.ByName("id")
	var title, content string
	err := c.QueryRow("SELECT title, content FROM posts WHERE rowid = ?", postId).Scan(&title, &content)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	post := NewPost(title, content)
	post.Id = postId // manually set post id
	c.HTML(rw, http.StatusOK, "edit", post)
	return nil
}

func (c *PostsController) Update(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	postId, title, content := p.ByName("id"), r.FormValue("title"), r.FormValue("content")
	_, err := c.Exec("UPDATE posts SET title = ?, content = ? WHERE rowid = ?", title, content, postId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	http.Redirect(rw, r, fmt.Sprintf("/posts/%s", postId), http.StatusSeeOther)
	return nil
}

// override Action to 404 page
func (c *PostsController) Action(a Action) httprouter.Handle {
	return httprouter.Handle(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := a(rw, r, p); err != nil {
			http.ServeFile(rw, r, path.Join("public", "404.html"))
		}
	})
}
