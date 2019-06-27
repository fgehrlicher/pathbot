package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type PathbotDirection struct {
	Direction string `json:"direction"`
}

func main() {
	grid := NewGrid()

	tile := Tile{}
	tile.Active = true
	tile.Exits.North = true
	tile.MazeExitDistance = 31
	tile.MazeExitDirection = "N"

	grid.Grid = append(grid.Grid, make([]Tile, 0))
	grid.Grid[0] = append(grid.Grid[0], tile)
	grid.Render()

	//var location = start()
	//explore(location)
}

func start() Tile {
	return apiPost("/pathbot/start", strings.NewReader("{}"))
}

func explore(location Tile) {
	reader := bufio.NewReader(os.Stdin)

	for {
		printLocation(location)

		if location.Status == "finished" {
			fmt.Println(location.Message)
			os.Exit(0)
		}

		direction, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		dir := PathbotDirection{Direction: strings.ToUpper(direction[0:1])}

		body, err := json.Marshal(dir)

		if err != nil {
			panic(err.Error())
		}

		location = apiPost(location.LocationPath, bytes.NewBuffer(body))
	}
}

func printLocation(location Tile) {
	fmt.Println()
	fmt.Println(location.Message)
	fmt.Println(location.Description)
}

func apiPost(path string, body io.Reader) Tile {
	domain := "https://api.noopschallenge.com"
	res, err := http.Post(domain+path, "application/json", body)
	if err != nil {
		panic(err.Error())
	}

	return parseResponse(res)
}

func parseResponse(res *http.Response) Tile {
	var response Tile
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err.Error())
	}

	return response
}
