package pkg

import (
	"bufio"
	"io"
)

type DiagnosticReport [][]byte

func ParseDiagnosticInput(input io.Reader) DiagnosticReport {
	result := make(DiagnosticReport, 0, 500)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		result = append(result, parseDiagnosticLine(scanner.Text()))
	}

	return result
}

func parseDiagnosticLine(line string) []byte {
	result := make([]byte, 12)

	for i, b := range line {
		if string(b) == "1" {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}
