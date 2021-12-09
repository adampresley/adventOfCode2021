package main

import "sync"

type FleetManager struct {
	submarines CrabSubmarineCollection
	wg         *sync.WaitGroup
	ch         chan Winner
}

func NewFleetManager(submarines CrabSubmarineCollection, wg *sync.WaitGroup, ch chan Winner) *FleetManager {
	return &FleetManager{
		submarines: submarines,
		wg:         wg,
		ch:         ch,
	}
}

func (fm *FleetManager) CalculateFuelCostToMoveTo(position int) {
	fm.wg.Add(1)

	go func() {
		result := 0

		for _, s := range fm.submarines {
			result += s.CalculateFuelCostToMove(position)
		}

		fm.ch <- Winner{
			FuelCost:           result,
			HorizontalPosition: position,
		}

		fm.wg.Done()
	}()
}
