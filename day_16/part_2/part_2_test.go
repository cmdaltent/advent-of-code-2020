package part_2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmdaltent/advent-of-code-2020/day_16/common"
)

func Test_discardInvalidNearbyTickets(t *testing.T) {
	fields := map[string]common.Constraint{
		"class": {common.Interval{Min: 1, Max: 3}, common.Interval{Min: 5, Max: 7}},
		"row":   {common.Interval{Min: 6, Max: 11}, common.Interval{Min: 33, Max: 44}},
		"seat":  {common.Interval{Min: 13, Max: 40}, common.Interval{Min: 45, Max: 50}},
	}

	nearbyTickets := [][]int{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	validTickets := [][]int{
		{7, 3, 47},
	}

	assert.Equal(t, validTickets, discardInvalidNearbyTickets(fields, nearbyTickets))
}
