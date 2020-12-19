package part_1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmdaltent/advent-of-code-2020/day_12/common"
)

func TestHeading_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		head     common.Heading
		degree   int
		expected common.Heading
	}{
		{name: "North 360 North", head: common.North, degree: 360, expected: common.North},
		{name: "South 360 South", head: common.South, degree: 360, expected: common.South},
		{name: "North 90 West", head: common.North, degree: 90, expected: common.West},
		{name: "West 270 North", head: common.West, degree: 270, expected: common.North},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.head.TurnLeft(test.degree))
		})
	}
}

func TestHeading_TurnRight(t *testing.T) {
	tests := []struct {
		name     string
		head     common.Heading
		degree   int
		expected common.Heading
	}{
		{name: "North 360 North", head: common.North, degree: 360, expected: common.North},
		{name: "South 360 South", head: common.South, degree: 360, expected: common.South},
		{name: "North 90 East", head: common.North, degree: 90, expected: common.East},
		{name: "West 270 South", head: common.West, degree: 270, expected: common.South},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.head.TurnRight(test.degree))
		})
	}
}
