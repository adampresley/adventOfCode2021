package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/adampresley/advent2021/cmd/day3/part2/internal/co2scrubber"
	"github.com/adampresley/advent2021/cmd/day3/part2/internal/lifesupport"
	"github.com/adampresley/advent2021/cmd/day3/part2/internal/oxygengenerator"
	"github.com/adampresley/advent2021/cmd/day3/pkg"
)

func main() {
	var (
		err      error
		rawBytes []byte
	)

	/*
	 * Open file then parse into a slice of byte arrays
	 */
	if rawBytes, err = ioutil.ReadFile("../../../inputs/day3.txt"); err != nil {
		panic(err)
	}

	/*
	 * Calculate oxygen generator rating and co2 scrubber rating, then
	 * calculate the life support rating.
	 */
	o2Rating := oxygengenerator.ProcessDiagnosticReport(pkg.ParseDiagnosticInput(bytes.NewReader(rawBytes)))
	co2Rating := co2scrubber.ProcessDiagnosticReport(pkg.ParseDiagnosticInput(bytes.NewReader(rawBytes)))
	lifeSupportRating := lifesupport.CalculateLifeSupportRating(o2Rating, co2Rating)

	fmt.Printf("O2 Rating: %d\n", o2Rating)
	fmt.Printf("CO2: %d\n", co2Rating)
	fmt.Printf("Life Support Rating: %d\n", lifeSupportRating)
}
