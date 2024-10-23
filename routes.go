package main

import (
	"net/http"
)

func buildRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /room/{id}", getRoom)
	mux.HandleFunc("POST /room/{id}", updateRoom)

	mux.Handle("/", http.FileServer(http.Dir("./static")))

	return mux
}
