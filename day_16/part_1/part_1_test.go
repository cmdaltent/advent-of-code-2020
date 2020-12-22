package part_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cmdaltent/advent-of-code-2020/day_16/input"
)

func TestPartOne(t *testing.T) {
	fields, _, nearbyTickets, err := input.ReadInput("testdata/input.txt")
	require.NoError(t, err)
	assert.Equal(t, 71, PartOne(fields, nearbyTickets))
}
