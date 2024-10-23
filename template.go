package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	_ "embed"

	"github.com/valyala/fasttemplate"
)

//go:embed tpl/page.html
var gridPage string
var gridPageTemplate *fasttemplate.Template

//go:embed tpl/row.html
var gridRow string
var gridRowTemplate *fasttemplate.Template

//go:embed tpl/cell.html
var gridCell string
var gridCellTemplate *fasttemplate.Template

func init() {
	gridPageTemplate = fasttemplate.New(gridPage, "{{", "}}")
	gridRowTemplate = fasttemplate.New(gridRow, "{{", "}}")
	gridCellTemplate = fasttemplate.New(gridCell, "{{", "}}")
}

func generateGridRow(roomId int, y int) (r string) {
	cells := ""
	for x := 0; x < gridWidth; x++ {
		v := rooms[roomId][x][y]

		cells += gridCellTemplate.ExecuteString(map[string]any{
			"roomId": strconv.Itoa(roomId),
			"x":      strconv.Itoa(x),
			"y":      strconv.Itoa(y),
			"v":      v,
		})
	}

	return gridRowTemplate.ExecuteString(map[string]any{
		"cells": cells,
	})
}

func generatePage(roomId int) (p string) {
	grid := ""
	for y := 0; y < gridHeight; y++ {
		r := generateGridRow(roomId, y)
		grid += r
	}

	p = gridPageTemplate.ExecuteString(map[string]any{
		"grid": grid,
	})

	return p
}

func generateRandomHex() string {
	r := rand.IntN(256)
	g := rand.IntN(256)
	b := rand.IntN(256)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func createNewGrid() grid {
	g := grid{}

	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			g[i][j] = generateRandomHex()
		}
	}

	return g
}
