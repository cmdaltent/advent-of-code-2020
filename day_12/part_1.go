package main

import "math"

func partOne(navigationInstructions []navigationInstruction) int {
	s := newShip(east)

	for _, instruction := range navigationInstructions {
		s.Apply(instruction)
	}

	return int(math.Abs(float64(s.pos.x)) + math.Abs(float64(s.pos.y)))
}
