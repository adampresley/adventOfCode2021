package decider

import (
	"github.com/adampresley/advent2021/cmd/day3/pkg"
)

func Decide(input pkg.DiagnosticReport, whichToKeep func(bit0Count, bit1Count int) byte) string {
	var (
		currentBit byte
		bit        byte
	)

	for bitPosition := 0; bitPosition < len(input[0]); bitPosition++ {
		bit0Count := 0
		bit1Count := 0

		if len(input) <= 1 {
			break
		}

		for y := 0; y < len(input); y++ {
			currentBit = input[y][bitPosition]
			if currentBit == 0 {
				bit0Count++
			} else {
				bit1Count++
			}
		}

		// If this returns false, track which rows we'll cut that
		// don't have this bit in the current bit position.
		if whichToKeep(bit0Count, bit1Count) == 0 {
			bit = 0
		} else {
			bit = 1
		}

		for y := 0; y < len(input); y++ {
			if input[y][bitPosition] != bit {
				input = append(input[:y], input[y+1:]...)
				y--
			}
		}
	}

	return byteArrayToString(input[0])
}

func byteArrayToString(input []byte) string {
	result := ""

	for _, b := range input {
		if b == 1 {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result
}
