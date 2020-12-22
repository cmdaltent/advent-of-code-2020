package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/day_16/input"
	"github.com/cmdaltent/advent-of-code-2020/day_16/part_1"
	"github.com/cmdaltent/advent-of-code-2020/day_16/part_2"
)

func main() {
	fields, myTicket, nearbyTickets, err := input.ReadInput("day_16/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("Part One - Ticket Error Rate: %d\n", part_1.PartOne(fields, nearbyTickets))
	log.Printf("Part Two - Departure Multiplied: %d\n", part_2.PartTwo(fields, myTicket, nearbyTickets))
}
