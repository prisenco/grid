package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getIdFromRequest(r *http.Request) (int, error) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)

	return id, err
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)

	if err != nil {
		fmt.Fprintf(w, "Error, id is not an integer")
		return
	}

	if id > 9 {
		fmt.Fprintf(w, "Choose a room between 1-9")
		return
	}

	if rooms[id] == nil {
		g := createNewGrid()
		rooms[id] = &g
	}

	gridTemplate := generateGrid(id)

	fmt.Fprint(w, gridTemplate)
}

func updateRoom(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)

	if err != nil {
		log.Fatal("Error, id is not an integer")
	}

	// Error if the room doesn't exist
	if rooms[id] == nil {
		log.Fatal("Room does not exist")
	}

	err = r.ParseForm()

	if err != nil {
		log.Fatal("Could not parse form")
	}

	xStr := r.FormValue("x")
	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Fatal("Could not get x value")
	}

	yStr := r.FormValue("y")
	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Fatal("Could not get y value")
	}

	color := r.FormValue("color")

	rooms[id][x][y] = color

	g := generateGrid(id)

	fmt.Fprint(w, g)
}
