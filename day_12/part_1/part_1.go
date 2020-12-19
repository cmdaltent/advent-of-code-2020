package part_1

import (
	"math"

	"github.com/cmdaltent/advent-of-code-2020/day_12/common"
)

func PartOne(navigationInstructions []common.NavigationInstruction) int {
	s := newShip(common.East)

	for _, instruction := range navigationInstructions {
		s.Apply(instruction)
	}

	return int(math.Abs(float64(s.position.X)) + math.Abs(float64(s.position.Y)))
}
