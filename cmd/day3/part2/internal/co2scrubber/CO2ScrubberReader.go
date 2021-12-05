package co2scrubber

import (
	"fmt"
	"strconv"

	"github.com/adampresley/advent2021/cmd/day3/part2/internal/decider"
	"github.com/adampresley/advent2021/cmd/day3/pkg"
)

func ProcessDiagnosticReport(input pkg.DiagnosticReport) (rating int) {
	fmt.Printf("Reading CO2 Generator Diagnostics...\n")

	result := decider.Decide(input, func(bit0Count, bit1Count int) byte {
		if bit0Count > bit1Count {
			return 1
		} else {
			return 0
		}
	})

	o2, _ := strconv.ParseUint(result, 2, 64)
	rating = int(o2)
	return
}
