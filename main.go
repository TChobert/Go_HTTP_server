package main

import (
	"fmt"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test reached\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
