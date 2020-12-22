package part_1

import "github.com/cmdaltent/advent-of-code-2020/day_16/common"

func PartOne(fields map[string][]common.Interval, nearbyTickets [][]int) int {
	var sum int
	for _, ticket := range nearbyTickets {
	ticketLoop:
		for _, v := range ticket {
			var valid bool
			for _, intervals := range fields {
				for _, interval := range intervals {
					valid = valid || interval.IsIn(v)
				}
				if valid {
					continue ticketLoop
				}
			}
			sum += v
		}
	}

	return sum
}
