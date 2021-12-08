package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fish []int

func main() {
	var (
		err    error
		bytes  []byte
		fishes Fish
	)

	bytes, err = ioutil.ReadFile("../../inputs/day6.txt")
	// bytes, err = ioutil.ReadFile("../../inputs/day6-sample.txt")

	if err != nil {
		panic(err)
	}

	/*****************************************************************************
	 * Part 1
	 ****************************************************************************/

	fmt.Printf("PART 1:\n-----------------------------\n\n")

	// Parse the initial age list
	fishes = buildFishList(strings.NewReader(string(bytes)))
	fmt.Printf("Initial age list:\n%v\n", fishes)

	for iteration := 0; iteration < 80; iteration++ {
		l := len(fishes)

		for index := 0; index < l; index++ {
			if fishes[index] == 0 {
				fishes = append(fishes, 8)
				fishes[index] = 7
			}

			fishes[index]--
		}

		// fmt.Printf("%v\n", fishes)
	}

	fmt.Printf("Num fishes: %d\n", len(fishes))

	/*****************************************************************************
	 * Part 2
	 ****************************************************************************/

	fmt.Printf("PART 2:\n-----------------------------\n\n")

	// Parse the initial age list
	fishes = buildFishList2(strings.NewReader(string(bytes)))
	fmt.Printf("Initial state: %v\n", fishes)

	for iteration := 0; iteration < 256; iteration++ {
		fishes = append(fishes, fishes[0])
		fishes = append(fishes[:0], fishes[1:]...)
		fishes[6] += fishes[8]
	}

	fmt.Printf("Num fishes: %d\n", sum(fishes))

}

func buildFishList(reader io.Reader) (result Fish) {
	b, _ := ioutil.ReadAll(reader)
	ageListString := string(b)
	split := strings.Split(strings.TrimSpace(ageListString), ",")

	for _, a := range split {
		age, _ := strconv.Atoi(strings.TrimSpace(a))
		result = append(result, age)
	}

	return
}

// Oooo we treat is like a state machine! Each array index is a count of how many fish
// we have for each count. For example, if we have 4 fish at a time of 3, position 3
// in the array will have a value of 4.
func buildFishList2(reader io.Reader) (result Fish) {
	b, _ := ioutil.ReadAll(reader)
	ageListString := string(b)
	split := strings.Split(strings.TrimSpace(ageListString), ",")

	for i := 0; i < 9; i++ {
		result = append(result, 0)
	}

	for _, a := range split {
		age, _ := strconv.Atoi(strings.TrimSpace(a))
		result[age]++
	}

	return
}

func sum(fishes Fish) (result int) {
	for _, a := range fishes {
		result += a
	}

	return
}
