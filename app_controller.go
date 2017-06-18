package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Action func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error

type AppController struct{}

// helps with error handling
func (c *AppController) Action(a Action) httprouter.Handle {
	return httprouter.Handle(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := a(rw, r, p); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
