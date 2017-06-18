package main

import (
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/markdown", http.HandlerFunc(GenerateMarkdown))
	http.Handle("/", http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	rw.Write(markdown)
}
