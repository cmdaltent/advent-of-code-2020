package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/day_16/input"
	"github.com/cmdaltent/advent-of-code-2020/day_16/part_1"
)

func main() {
	fields, _, nearbyTickets, err := input.ReadInput("day_16/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("Part One - Ticket Error Rate: %d\n", part_1.PartOne(fields, nearbyTickets))
}
