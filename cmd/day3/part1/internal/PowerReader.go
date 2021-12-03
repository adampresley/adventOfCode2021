package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adampresley/advent2021/cmd/day3/pkg"
)

func CalculateGammaAndEpsilon(diagnosticReport pkg.DiagnosticReport) (gamma, epsilon int) {
	var (
		lastBitPos    int
		bit           byte
		bitOneCount   int
		bitZeroCount  int
		gammaString   strings.Builder
		epsilonString strings.Builder
	)

	lastBitPos = len(diagnosticReport[0]) - 1

	for x := 0; x <= lastBitPos; x++ {
		bitOneCount = 0
		bitZeroCount = 0

		for y, _ := range diagnosticReport {
			bit = diagnosticReport[y][x]

			if bit == 0 {
				bitZeroCount++
			} else {
				bitOneCount++
			}
		}

		if bitOneCount > bitZeroCount {
			gammaString.WriteString("1")
			epsilonString.WriteString("0")
		} else {
			gammaString.WriteString("0")
			epsilonString.WriteString("1")
		}
	}

	fmt.Printf("Gamma String: %s\n", gammaString.String())
	fmt.Printf("Epsilon String: %s\n", epsilonString.String())
	fmt.Printf("\n")

	g, _ := strconv.ParseUint(gammaString.String(), 2, 64)
	e, _ := strconv.ParseUint(epsilonString.String(), 2, 64)

	gamma = int(g)
	epsilon = int(e)
	return
}

func CalculatePowerConsumption(gamma, epsilon int) int {
	return gamma * epsilon
}
