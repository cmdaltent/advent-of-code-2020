package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/day_12/input"
)

func main() {
	navigationInstructions, err := input.ReadNavigationInstructions("day_12/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("Part One - Manhatten Distance: %d", PartOne(navigationInstructions))
	log.Printf("Part Two - Manhatten Distance: %d", PartTwo(navigationInstructions))
}
