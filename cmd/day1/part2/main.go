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
		err    error
		input  []byte
		result int
	)

	if input, err = os.ReadFile("../../../inputs/day1.txt"); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	values := makeArray(scanner)

	for i := 0; i < len(values)-3; i++ {
		if values[i]+values[i+1]+values[i+2] < values[i+1]+values[i+2]+values[i+3] {
			result++
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func makeArray(scanner *bufio.Scanner) []int {
	var (
		result []int
	)

	result = make([]int, 0, 500)

	for scanner.Scan() {
		value := getValue(scanner.Text())
		result = append(result, value)
	}

	return result
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
