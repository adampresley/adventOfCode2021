package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var (
		err   error
		bytes []byte
	)

	if bytes, err = ioutil.ReadFile("../../inputs/day5.txt"); err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(bytes))
	inputReader := strings.NewReader(input)

	/*****************************************************************************
	 * Part 1
	 ****************************************************************************/
	fmt.Printf("\nPart 1:\n--------------------------------------------------\n")

	grid := NewGrid(inputReader, false)

	fmt.Printf("\n")

	overlaps := grid.CountOverlaps()
	fmt.Printf("Result: %d\n", overlaps)

	fmt.Printf("\n")

	/*****************************************************************************
	 * Part 2
	 ****************************************************************************/
	fmt.Printf("\nPart 2:\n--------------------------------------------------\n")
	inputReader = strings.NewReader(input)
	grid = NewGrid(inputReader, true)

	fmt.Printf("\n")

	overlaps = grid.CountOverlaps()
	fmt.Printf("Result: %d\n", overlaps)

	fmt.Printf("\n")

}

func printLines(lines LineCollection) {
	fmt.Printf("Lines:\n--------------------------------\n")
	fmt.Printf("%s\n", lines.String())
}

func printGrid(grid *Grid) {
	fmt.Printf("Grid:\n--------------------------------\n")
	fmt.Printf("%s\n", grid.String())
}
