package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	tests := []struct {
		name             string
		startingNumbers  []int
		terminationRound int
		expectedResult   int
	}{
		{name: "2020th of [0,3,6]", startingNumbers: []int{0, 3, 6}, terminationRound: 2020, expectedResult: 436},
		{name: "2020th of [1,3,2]", startingNumbers: []int{1, 3, 2}, terminationRound: 2020, expectedResult: 1},
		{name: "2020th of [2,1,3]", startingNumbers: []int{2, 1, 3}, terminationRound: 2020, expectedResult: 10},
		{name: "2020th of [1,2,3]", startingNumbers: []int{1, 2, 3}, terminationRound: 2020, expectedResult: 27},
		{name: "2020th of [2,3,1]", startingNumbers: []int{2, 3, 1}, terminationRound: 2020, expectedResult: 78},
		{name: "2020th of [3,2,1]", startingNumbers: []int{3, 2, 1}, terminationRound: 2020, expectedResult: 438},
		{name: "2020th of [3,1,2]", startingNumbers: []int{3, 1, 2}, terminationRound: 2020, expectedResult: 1836},

		{name: "30000000th of [0,3,6]", startingNumbers: []int{0, 3, 6}, terminationRound: 30000000, expectedResult: 175594},
		{name: "30000000th of [1,3,2]", startingNumbers: []int{1, 3, 2}, terminationRound: 30000000, expectedResult: 2578},
		{name: "30000000th of [2,1,3]", startingNumbers: []int{2, 1, 3}, terminationRound: 30000000, expectedResult: 3544142},
		{name: "30000000th of [1,2,3]", startingNumbers: []int{1, 2, 3}, terminationRound: 30000000, expectedResult: 261214},
		{name: "30000000th of [2,3,1]", startingNumbers: []int{2, 3, 1}, terminationRound: 30000000, expectedResult: 6895259},
		{name: "30000000th of [3,2,1]", startingNumbers: []int{3, 2, 1}, terminationRound: 30000000, expectedResult: 18},
		{name: "30000000th of [3,1,2]", startingNumbers: []int{3, 1, 2}, terminationRound: 30000000, expectedResult: 362},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedResult, play(test.startingNumbers, test.terminationRound))
		})
	}
}
