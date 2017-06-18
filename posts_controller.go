package main

import (
	"database/sql"
	"errors"
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
	var title, content string
	err := c.QueryRow("select title, content from posts").Scan(&title, &content)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	c.Data(rw, 200, []byte(fmt.Sprintf("title is %s and content is %s", title, content)))
	return nil
}

func (c *PostsController) Show(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	return errors.New("Damn")
}

// override Action to 404 page
func (c *PostsController) Action(a Action) httprouter.Handle {
	return httprouter.Handle(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := a(rw, r, p); err != nil {
			http.ServeFile(rw, r, path.Join("public", "404.html"))
		}
	})
}
