package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cmdaltent/advent-of-code-2020/day_12/input"
)

func Test_PartOne(t *testing.T) {
	navigationInstructions, err := input.ReadNavigationInstructions("testdata/partOne_test_input.txt")
	require.NoError(t, err)

	assert.Equal(t, 25, PartOne(navigationInstructions))
}

func Test_PartTwo(t *testing.T) {
	navigationInstructions, err := input.ReadNavigationInstructions("testdata/partOne_test_input.txt")
	require.NoError(t, err)

	assert.Equal(t, 286, PartTwo(navigationInstructions))
}
