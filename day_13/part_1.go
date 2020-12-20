package main

import "math"

func partOne(earliestDeparture int, busLines []int) int {
	earliestBusDeparture := math.MaxInt64
	var earliestBus int
	for _, bus := range busLines {
		if depart := ((earliestDeparture / bus) + 1) * bus; depart < earliestBusDeparture {
			earliestBusDeparture = depart
			earliestBus = bus
		}
	}
	return earliestBus * (earliestBusDeparture - earliestDeparture)
}
