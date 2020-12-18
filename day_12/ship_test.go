package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		head     heading
		degree   int
		expected heading
	}{
		{name: "North 360 North", head: north, degree: 360, expected: north},
		{name: "South 360 South", head: south, degree: 360, expected: south},
		{name: "North 90 West", head: north, degree: 90, expected: west},
		{name: "West 270 North", head: west, degree: 270, expected: north},
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
		head     heading
		degree   int
		expected heading
	}{
		{name: "North 360 North", head: north, degree: 360, expected: north},
		{name: "South 360 South", head: south, degree: 360, expected: south},
		{name: "North 90 East", head: north, degree: 90, expected: east},
		{name: "West 270 South", head: west, degree: 270, expected: south},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.head.TurnRight(test.degree))
		})
	}
}
