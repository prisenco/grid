package main

import (
	"log"
	"net/http"

	_ "embed"
)

const gridHeight = 3
const gridWidth = 3

type grid [gridWidth][gridHeight]string

var rooms map[int]*grid

func init() {
	rooms = make(map[int]*grid, 9)
}

func main() {
	mux := buildRoutes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
