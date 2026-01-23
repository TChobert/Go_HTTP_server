package router

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go server with chi!\n")
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from test location!\n")
	})
	return r
}
