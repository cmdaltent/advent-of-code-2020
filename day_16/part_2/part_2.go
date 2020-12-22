package part_2

import (
	"sort"
	"strings"

	"github.com/cmdaltent/advent-of-code-2020/day_16/common"
)

func PartTwo(fields map[string]common.Constraint, myTicket []int, nearbyTickets [][]int) int {
	validNearbyTickets := discardInvalidNearbyTickets(fields, nearbyTickets)

	// backtracking in ascending order works best if input is sorted
	ticketFieldOrder := ticketFieldOrderAscending(fields, validNearbyTickets)

	// backtracking
	found, mapping := fieldAssignment(fields, validNearbyTickets, ticketFieldOrder, make(map[string]int))
	if !found {
		return -1 // hack. I am to lazy to introduce an error
	}

	departureMultiplied := 1
	for field, pos := range mapping {
		if strings.HasPrefix(field, "departure") {
			departureMultiplied *= myTicket[pos]
		}
	}

	return departureMultiplied
}

func discardInvalidNearbyTickets(fields map[string]common.Constraint, nearbyTickets [][]int) [][]int {
	var invalidTicketIndices []int
	for i, ticket := range nearbyTickets {
		for _, v := range ticket {
			var valid bool
			for _, constraint := range fields {
				valid = valid || constraint.Satisfies(v)
			}
			if !valid {
				invalidTicketIndices = append(invalidTicketIndices, i)
			}
		}
	}

	var validTickets [][]int
outerLoop:
	for i, ticket := range nearbyTickets {
		for j := 0; j < len(invalidTicketIndices); j++ {
			if i == invalidTicketIndices[j] {
				continue outerLoop
			}
		}
		validTickets = append(validTickets, ticket)
	}

	return validTickets
}

type result struct {
	ticketFieldIndex int
	count            int
}

func ticketFieldOrderAscending(fields map[string]common.Constraint, validNearbyTickets [][]int) []*result {
	var order []*result
	for tf := 0; tf < len(fields); tf++ {
		var count int
		for _, constraint := range fields {
			valid := true
			for _, ticket := range validNearbyTickets {
				if !constraint.Satisfies(ticket[tf]) {
					valid = false
				}
			}
			if valid {
				count++
			}
		}
		order = append(order, &result{ticketFieldIndex: tf, count: count})
	}
	sort.Slice(order, func(i, j int) bool {
		return order[i].count < order[j].count
	})
	return order
}

func fieldAssignment(fields map[string]common.Constraint, validNearbyTickets [][]int, ticketFieldOrder []*result, skipFields map[string]int) (bool, map[string]int) {
	if len(skipFields) == len(fields) {
		return true, skipFields
	}

ticketFieldLoop:
	for _, ticketField := range ticketFieldOrder {
		for _, tf := range skipFields {
			if ticketField.ticketFieldIndex == tf {
				continue ticketFieldLoop
			}
		}
	fieldLoop:
		for field, constraint := range fields {
			if _, ok := skipFields[field]; ok {
				continue
			}
			for _, ticket := range validNearbyTickets {
				if !constraint.Satisfies(ticket[ticketField.ticketFieldIndex]) {
					continue fieldLoop
				}
			}
			cpySF := make(map[string]int, len(skipFields))
			for k, v := range skipFields {
				cpySF[k] = v
			}
			cpySF[field] = ticketField.ticketFieldIndex
			if found, mapping := fieldAssignment(fields, validNearbyTickets, ticketFieldOrder, cpySF); found {
				return true, mapping
			}
		}
	}

	return false, nil
}
