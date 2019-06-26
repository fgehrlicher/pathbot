package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
)

func NewGrid() *Grid {
	grid := Grid{}
	grid.Border.LeftUpperCorner = "╔"
	grid.Border.LeftBottomCorner = "╚"
	grid.Border.RightUpperCorner = "╗"
	grid.Border.RightBottomCorner = "╝"
	grid.Border.HorizontalSideBorder = "║"
	grid.Border.VerticalSideBorder = "═"
	return &grid
}

type Grid struct {
	Grid   [][]Tile
	Border struct {
		LeftUpperCorner      string
		LeftBottomCorner     string
		RightUpperCorner     string
		RightBottomCorner    string
		HorizontalSideBorder string
		VerticalSideBorder   string
	}
}

type Tile struct {
	Active             bool
	UnderlyingLocation PathbotLocation
	Exits              struct {
		North bool
		East  bool
		South bool
		West  bool
	}
	MazeExitDirection string
	MazeExitDistance  int
	LocationPath      string
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

			titleRowFilling := strings.Repeat(grid.Border.VerticalSideBorder, width-2)
			bottomRowFilling := titleRowFilling
			emptyRowFilling := strings.Repeat(" ", width-2)
			rightBorder := grid.Border.HorizontalSideBorder
			leftBorder := grid.Border.HorizontalSideBorder

			distanceString := fmt.Sprintf(" Distance: %v ", tile.MazeExitDistance)
			directionString := fmt.Sprintf(" Direction: %v ", tile.MazeExitDirection)
			distanceRowFilling := distanceString + strings.Repeat(" ", (width-2)-len(distanceString))
			directionRowFilling := directionString + strings.Repeat(" ", (width-2)-len(directionString))

			if tile.Exits.North {
				titleRowFilling = aurora.Green(titleRowFilling).String()
			}
			if tile.Exits.South {
				bottomRowFilling = aurora.Green(bottomRowFilling).String()
			}
			if tile.Exits.West {
				leftBorder = aurora.Green(leftBorder).String()
			}
			if tile.Exits.East {
				rightBorder = aurora.Green(rightBorder).String()
			}

			renderMap[currentRowStart] += grid.Border.LeftUpperCorner + titleRowFilling + grid.Border.RightUpperCorner
			renderMap[currentRowStart+1] += leftBorder + emptyRowFilling + rightBorder
			renderMap[currentRowStart+2] += leftBorder + distanceRowFilling + rightBorder
			renderMap[currentRowStart+3] += leftBorder + directionRowFilling + rightBorder
			renderMap[currentRowStart+4] += leftBorder + emptyRowFilling + rightBorder
			renderMap[currentRowStart+5] += grid.Border.LeftBottomCorner + bottomRowFilling + grid.Border.RightBottomCorner
		}
		currentRowStart = currentRowStart + height
	}

	for i := 0; i < len(renderMap); i ++ {
		fmt.Println(renderMap[i])
	}
}
