package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/day_12/input"
	"github.com/cmdaltent/advent-of-code-2020/day_12/part_1"
)

func main() {
	navigationInstructions, err := input.ReadNavigationInstructions("day_12/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("Part One - Manhatten Distance: %d", part_1.PartOne(navigationInstructions))
}
