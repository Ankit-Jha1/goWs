package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting web server")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		// log.Fatal ==> exits with 1
		log.Fatal("web server failed")
	}
}
