package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

func main() {
	db := NewDB()

	c := &PostsController{
		Render: render.New(render.Options{}),
		DB:     db,
	}

	router := httprouter.New()
	router.GET("/posts", c.Action(c.Index))
	router.GET("/posts/:id", c.Action(c.Show))
	router.GET("/new_post", c.Action(c.New))
	router.POST("/posts", c.Action(c.Create))
	router.GET("/posts/:id/edit", c.Action(c.Edit))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", router)
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists posts(title text, content text)")
	if err != nil {
		panic(err)
	}
	return db
}
