package part_1

import "github.com/cmdaltent/advent-of-code-2020/day_16/common"

func PartOne(fields map[string]common.Constraint, nearbyTickets [][]int) int {
	var sum int
	for _, ticket := range nearbyTickets {
	ticketLoop:
		for _, v := range ticket {
			for _, constraint := range fields {
				if constraint.Satisfies(v) {
					continue ticketLoop
				}
			}
			sum += v
		}
	}

	return sum
}
