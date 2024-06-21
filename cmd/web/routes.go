package main

import (
	"net/http"

	"github.com/Ankit-Jha1/goWs/internal/handlers"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	// dont put () in front of the function
	mux.HandleFunc("GET /", handlers.Home)

	return mux
}
