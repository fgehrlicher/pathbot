package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
)

func NewGrid() *Grid {
	return &Grid{}
}

type Grid struct {
	Grid [][]Tile
}

type Tile struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Active      bool
	RawExits    []string `json:"exits"`
	Exits       struct {
		North bool
		East  bool
		South bool
		West  bool
	}
	MazeExitDirection string `json:"mazeExitDirection"`
	MazeExitDistance  int    `json:"mazeExitDistance"`
	LocationPath      string `json:"locationPath"`
}

func (grid Grid) Render() {
	currentRowStart := 0
	width := 16
	height := 6

	renderMap := make([]string, height)
	for row := 0; row < len(grid.Grid[0]); row ++ {
		renderMap = append(renderMap, make([]string, height)...)
		for column := 0; column < len(grid.Grid); column ++ {
			tile := grid.Grid[column][row]

			borderBlock := " "
			bodyBlock := " "
			titleRowFilling := strings.Repeat(bodyBlock, width-2)
			bottomRowFilling := titleRowFilling
			emptyRowFilling := strings.Repeat(bodyBlock, width-2)

			redBorder := aurora.BgBrightRed(borderBlock).String()
			leftBottomCorner := redBorder
			rightBottomCorner := redBorder
			rightUpperCorner := redBorder
			leftUpperCorner := redBorder
			leftBorder := redBorder
			rightBorder := redBorder

			distanceString := fmt.Sprintf(bodyBlock+"Distance: %v"+bodyBlock, tile.MazeExitDistance)
			directionString := fmt.Sprintf(bodyBlock+"Direction: %v"+bodyBlock, tile.MazeExitDirection)
			distanceRowFilling := distanceString + strings.Repeat(bodyBlock, (width-2)-len(distanceString))
			directionRowFilling := directionString + strings.Repeat(bodyBlock, (width-2)-len(directionString))

			if tile.Active {
				distanceRowFilling = aurora.White(aurora.BgGreen(distanceRowFilling).String()).String()
				directionRowFilling = aurora.White(aurora.BgGreen(directionRowFilling).String()).String()
				emptyRowFilling = aurora.BgGreen(emptyRowFilling).String()
			} else {
				distanceRowFilling = aurora.Black(aurora.BgWhite(distanceRowFilling).String()).String()
				directionRowFilling = aurora.Black(aurora.BgWhite(directionRowFilling).String()).String()
				emptyRowFilling = aurora.BgWhite(emptyRowFilling).String()
			}

			if tile.Exits.North {
				titleRowFilling = aurora.BgBrightWhite(titleRowFilling).String()
			} else {
				titleRowFilling = aurora.BgBrightRed(titleRowFilling).String()
			}
			if tile.Exits.South {
				bottomRowFilling = aurora.BgBrightWhite(bottomRowFilling).String()
			} else {
				bottomRowFilling = aurora.BgBrightRed(bottomRowFilling).String()
			}
			if tile.Exits.West {
				leftBorder = aurora.BgBrightWhite(leftBorder).String()
			}
			if tile.Exits.East {
				rightBorder = aurora.BgBrightWhite(rightBorder).String()
			}

			renderMap[currentRowStart] += leftUpperCorner + titleRowFilling + rightUpperCorner
			renderMap[currentRowStart+1] += leftBorder + emptyRowFilling + rightBorder
			renderMap[currentRowStart+2] += leftBorder + distanceRowFilling + rightBorder
			renderMap[currentRowStart+3] += leftBorder + directionRowFilling + rightBorder
			renderMap[currentRowStart+4] += leftBorder + emptyRowFilling + rightBorder
			renderMap[currentRowStart+5] += leftBottomCorner + bottomRowFilling + rightBottomCorner
		}
		currentRowStart = currentRowStart + height
	}

	for i := 0; i < len(renderMap); i ++ {
		fmt.Println(renderMap[i])
	}
}
