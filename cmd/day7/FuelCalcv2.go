package main

import (
	"math"
)

func FuelCostCalcv2(initialPos, position int) (fuelCost int) {
	deltaString := math.Abs(float64(initialPos - position))
	delta := int(deltaString)

	for i := 1; i <= delta; i++ {
		fuelCost += i
	}

	// fmt.Printf("Move from %d to %d (delta %d): %d\n", initialPos, position, delta, fuelCost)
	return
}
