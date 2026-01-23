package server

import (
	"net/http"
	"time"
)

func NewServer(router http.Handler) *http.Server {

	return &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}
