package main

import (
	"fmt"
	"os"

	"github.com/adampresley/advent2021/cmd/day2/part2/internal"
)

func main() {
	var (
		err error
		fp  *os.File
	)

	/*
	 * Read input file
	 */
	if fp, err = os.Open("../../../inputs/day2.txt"); err != nil {
		panic(err)
	}

	/*
	 * Parse the file and move the sub!
	 */
	directionChanges := internal.Parse(fp)
	guidanceSystem := internal.NewGuidanceSystem()

	horizontalPos, depth := guidanceSystem.ProcessNavigationInput(directionChanges)
	fmt.Printf("Horizontal position: %d\nDepth: %d\n", horizontalPos, depth)
	fmt.Printf("Result: %d\n", horizontalPos*depth)
}
