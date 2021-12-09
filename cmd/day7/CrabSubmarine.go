package main

import (
	"strconv"
)

type CrabSubmarine struct {
	Position     int
	fuelCosts    map[int]int
	fuelCostFunc FuelCostAlgorithm
}

type CrabSubmarineCollection []*CrabSubmarine

func NewCrabSubmarine(fuel int, fuelCostFunc FuelCostAlgorithm) *CrabSubmarine {
	return &CrabSubmarine{
		Position:     fuel,
		fuelCosts:    make(map[int]int),
		fuelCostFunc: fuelCostFunc,
	}
}

func NewCrabSubmarineFromString(fuelString string, fuelCostFunc FuelCostAlgorithm) *CrabSubmarine {
	var (
		err  error
		fuel int
	)

	if fuel, err = strconv.Atoi(fuelString); err != nil {
		panic(err)
	}

	return NewCrabSubmarine(fuel, fuelCostFunc)
}

func (s *CrabSubmarine) CalculateFuelCostToMove(position int) (fuelCost int) {
	return s.fuelCostFunc(s.Position, position)
}
