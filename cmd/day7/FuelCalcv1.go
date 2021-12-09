package main

import (
	"math"
)

type FuelCostAlgorithm func(initialPos, position int) int

func FuelCostCalcv1(initialPos, position int) (fuelCost int) {
	result := math.Abs(float64(initialPos - position))
	fuelCost = int(result)

	// fmt.Printf("Move from %d to %d: %d\n", initialPos, position, fuelCost)
	return
}
