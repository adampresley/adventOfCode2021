package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		err           error
		input         []byte
		previousValue int
		currentValue  int
		result        int
	)

	if input, err = os.ReadFile("../../../inputs/day1.txt"); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	previousValue = 0

	for scanner.Scan() {
		value := scanner.Text()

		if previousValue == 0 {
			previousValue = getValue(value)
			continue
		}

		currentValue = getValue(value)

		if currentValue > previousValue {
			result++
		}

		previousValue = currentValue
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", result)
}

func getValue(textValue string) int {
	var (
		err    error
		result int
	)

	if result, err = strconv.Atoi(textValue); err != nil {
		panic(err)
	}

	return result
}
