package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

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
	mux := http.NewServeMux()

	mux.HandleFunc("GET /room/{id}", getRoom)
	mux.HandleFunc("POST /room/{id}", updateRoom)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func getIdFromRequest(r *http.Request) (int, error) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)

	return id, err
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)

	if err != nil {
		log.Fatal("Error, id is not an integer")
	}

	if rooms[id] == nil {
		g := createNewGrid()
		rooms[id] = &g
	}

	gridTemplate := generateGrid(id)

	fmt.Fprint(w, gridTemplate)
}

func generateGridRow(roomId int, y int) (r string) {
	r = "<div>"
	for x := 0; x < gridWidth; x++ {
		v := rooms[roomId][x][y]
		r = r + fmt.Sprintf(`<input onchange="updateCell(%d, %d, %d, this.value);" type="color" id="head" name="head" value="%s" />`, roomId, x, y, v)
	}
	r += "</div>"
	return r
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
	<form method="POST" target="/room/` + fmt.Sprint(roomId) + `">
		<input type="hidden" name="x" />
		<input type="hidden" name="y" />
		<input type="hidden" name="color" />
	</form>
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
