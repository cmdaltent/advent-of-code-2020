package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/day_14/part_1"
	"github.com/cmdaltent/advent-of-code-2020/day_14/part_2"
)

func main() {
	log.Printf("Part One - Sum of all Values: %d\n", part_1.PartOne("day_14/input.txt"))
	log.Printf("Part Two - Sum of all values: %d\n", part_2.PartTwo("day_14/input.txt"))
}
