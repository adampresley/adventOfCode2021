package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Winner struct {
	FuelCost           int
	HorizontalPosition int
}

func main() {
	var (
		err error
		b   []byte
	)

	// if b, err = ioutil.ReadFile("../../inputs/day7-sample.txt"); err != nil {
	if b, err = ioutil.ReadFile("../../inputs/day7.txt"); err != nil {
		panic(err)
	}

	max := getMax(string(b))

	/*****************************************************************************
	 * Part 1
	 ****************************************************************************/
	fmt.Printf("PART 1:\n--------------------------------------\n")

	wg := &sync.WaitGroup{}
	ch := make(chan Winner, 5000)

	crabSubmarines := initializeCrabSubmarinesv1(string(b))
	fleetManager := NewFleetManager(crabSubmarines, wg, ch)

	currentWinner := &Winner{
		FuelCost:           math.MaxInt,
		HorizontalPosition: 0,
	}

	go func() {
		for potentialWinner := range ch {
			if potentialWinner.FuelCost < currentWinner.FuelCost {
				currentWinner.FuelCost = potentialWinner.FuelCost
				currentWinner.HorizontalPosition = potentialWinner.HorizontalPosition
			}
		}
	}()

	for pos := 0; pos < max; pos++ {
		fleetManager.CalculateFuelCostToMoveTo(pos)
	}

	wg.Wait()
	fmt.Printf("The best position to go to is %d with a fuel cost of %d\n\n", currentWinner.HorizontalPosition, currentWinner.FuelCost)

	/*****************************************************************************
	 * Part 2
	 ****************************************************************************/
	fmt.Printf("PART 2:\n--------------------------------------\n")

	wg = &sync.WaitGroup{}
	ch = make(chan Winner, 5000)
	currentWinner = &Winner{
		FuelCost:           math.MaxInt,
		HorizontalPosition: 0,
	}

	crabSubmarines = initializeCrabSubmarinesv2(string(b))
	fleetManager = NewFleetManager(crabSubmarines, wg, ch)

	go func() {
		for potentialWinner := range ch {
			if potentialWinner.FuelCost < currentWinner.FuelCost {
				currentWinner.FuelCost = potentialWinner.FuelCost
				currentWinner.HorizontalPosition = potentialWinner.HorizontalPosition
			}
		}
	}()

	for pos := 0; pos < max; pos++ {
		fleetManager.CalculateFuelCostToMoveTo(pos)
	}

	wg.Wait()
	fmt.Printf("The best position to go to is %d with a fuel cost of %d\n", currentWinner.HorizontalPosition, currentWinner.FuelCost)
}

func initializeCrabSubmarinesv1(input string) (result CrabSubmarineCollection) {
	split := strings.Split(strings.TrimSpace(input), ",")

	for _, fuelString := range split {
		result = append(result, NewCrabSubmarineFromString(fuelString, FuelCostCalcv1))
	}

	return
}

func initializeCrabSubmarinesv2(input string) (result CrabSubmarineCollection) {
	split := strings.Split(strings.TrimSpace(input), ",")

	for _, fuelString := range split {
		result = append(result, NewCrabSubmarineFromString(fuelString, FuelCostCalcv2))
	}

	return
}

func getMax(input string) (result int) {
	split := strings.Split(strings.TrimSpace(input), ",")

	for _, fuelString := range split {
		fuel, _ := strconv.Atoi(fuelString)

		if fuel > result {
			result = fuel
		}
	}

	return
}
