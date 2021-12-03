package main

import (
	"fmt"
	"os"

	"github.com/adampresley/advent2021/cmd/day3/part1/internal"
	"github.com/adampresley/advent2021/cmd/day3/pkg"
)

func main() {
	var (
		err            error
		fp             *os.File
		rawDiagnostics pkg.DiagnosticReport
	)

	/*
	 * Open file then parse into a slice of byte arrays
	 */
	if fp, err = os.Open("../../../inputs/day3.txt"); err != nil {
		panic(err)
	}

	rawDiagnostics = pkg.ParseDiagnosticInput(fp)

	fmt.Printf("Raw diagnostic report:\n")

	for _, d := range rawDiagnostics {
		fmt.Printf("%b\n", d)
	}

	fmt.Printf("\n")

	/*
	 * Calculate gamma and epsilon, then the power consumption
	 */
	gamma, epsilon := internal.CalculateGammaAndEpsilon(rawDiagnostics)
	powerConsumption := internal.CalculatePowerConsumption(gamma, epsilon)

	fmt.Printf("Gamma: %d\n", gamma)
	fmt.Printf("Epsilon: %d\n", epsilon)
	fmt.Printf("Power consumption: %d\n", powerConsumption)
}
