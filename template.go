package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	_ "embed"

	"github.com/valyala/fasttemplate"
)

//go:embed tpl/row.html
var gridRow string
var gridRowTemplate *fasttemplate.Template

//go:embed tpl/cell.html
var gridCell string
var gridCellTemplate *fasttemplate.Template

func init() {
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

func generateGrid(roomId int) (t string) {
	t = `
	<style>
		input[type="color"] {
			width: 50px;
			height: 50px;
			border: none;
			padding: 0;
			background-color: #fff;
		}
	</style>
	<script type="text/javascript">
		async function updateCell(room, x, y, color) {
		  const url = "/room/` + fmt.Sprint(roomId) + `";
			const params = new URLSearchParams();
			params.set("x", x);
			params.set("y", y);
			params.set("color", color);

			try {
				const response = await fetch(url, {
					body: params,
					method: "POST"
				});
				if (!response.ok) {
					console.log (response.status);
					return;
				}
				const newBody = await response.text();
				document.body.innerHTML = newBody;
			} catch (e) {
				console.error(e.message);
			}
		}
	</script>
	`
	for y := 0; y < gridHeight; y++ {
		r := generateGridRow(roomId, y)
		t += r
	}

	return t
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
