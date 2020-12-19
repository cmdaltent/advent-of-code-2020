package main

import (
	"math"

	"github.com/cmdaltent/advent-of-code-2020/day_12/common"
	"github.com/cmdaltent/advent-of-code-2020/day_12/part_1"
	"github.com/cmdaltent/advent-of-code-2020/day_12/part_2"
)

func PartOne(navigationInstructions []common.NavigationInstruction) int {
	s := part_1.NewShip(common.East)

	for _, instruction := range navigationInstructions {
		s.Apply(instruction)
	}

	return int(math.Abs(float64(s.Position.X)) + math.Abs(float64(s.Position.Y)))
}

func PartTwo(navigationInstructions []common.NavigationInstruction) int {
	s := part_2.NewShip(common.Position{Y: 1, X: 10})

	for _, instruction := range navigationInstructions {
		s.Apply(instruction)
	}

	return int(math.Abs(float64(s.Position.X)) + math.Abs(float64(s.Position.Y)))
}
