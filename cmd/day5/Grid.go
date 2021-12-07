package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"
)

type Grid struct {
	lines  LineCollection
	Points [1000][1000]int
}

func NewGrid(input io.Reader, considerDiagonals bool) *Grid {
	result := &Grid{
		lines: make(LineCollection, 0, 500),
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var x1, x2, y1, y2 int
		inputLine := scanner.Text()

		inputPointPairs := strings.Split(inputLine, " -> ")

		inputPoint1 := strings.Split(inputPointPairs[0], ",")
		inputPoint2 := strings.Split(inputPointPairs[1], ",")

		x1, _ = strconv.Atoi(inputPoint1[0])
		y1, _ = strconv.Atoi(inputPoint1[1])
		x2, _ = strconv.Atoi(inputPoint2[0])
		y2, _ = strconv.Atoi(inputPoint2[1])

		newLine := NewLine(x1, y1, x2, y2)
		if considerDiagonals {
			result.lines = append(result.lines, newLine)
		} else if newLine.IsStraightLine() {
			result.lines = append(result.lines, newLine)
		}
	}

	result.plotLines()
	return result
}

func (g *Grid) CountOverlaps() int {
	result := 0

	for y := 0; y < len(g.Points); y++ {
		for x := 0; x < len(g.Points); x++ {
			if g.Points[y][x] > 1 {
				result++
			}
		}
	}

	return result
}

func (g *Grid) plotLines() {
	for _, line := range g.lines {
		g.plotLine(line.Point1, line.Point2)
	}
}

func (g *Grid) plotLine(point1, point2 image.Point) {
	stepX := 1
	stepY := 1
	x := point1.X
	y := point1.Y

	if point1.X > point2.X {
		stepX = -1
	}

	if point1.X == point2.X {
		stepX = 0
	}

	if point1.Y > point2.Y {
		stepY = -1
	}

	if point1.Y == point2.Y {
		stepY = 0
	}

	// Horizontal
	if stepX == 0 { // Horizontal
		for y != point2.Y+stepY {
			g.Points[y][x]++
			y += stepY
		}
	} else if stepY == 0 { // Vertical
		for x != point2.X+stepX {
			g.Points[y][x]++
			x += stepX
		}
	} else {

		for y != point2.Y+stepY { // Diagonal
			for x != point2.X+stepX {
				g.Points[y][x]++

				x += stepX
				y += stepY
			}
		}
	}
}

func (g *Grid) String() string {
	result := strings.Builder{}

	for y := 0; y < len(g.Points); y++ {
		for x := 0; x < len(g.Points); x++ {
			if g.Points[y][x] == 0 {
				result.WriteString(". ")
			} else {
				result.WriteString(fmt.Sprintf("%d ", g.Points[y][x]))
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}
