package main

import (
	"log"
	"os"
	"github.com/TChobert/Go_HTTPS_server/internal/router"
	"github.com/TChobert/Go_HTTPS_server/internal/server"
)

func main() {

	rout := router.NewRouter()
	srv := server.NewServer(rout)

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	go func() {
		if err := server.RunTLS(srv, certPath, keyPath); err != nil {
			log.Fatal(err)
		}
	} ()
	select {}
}
