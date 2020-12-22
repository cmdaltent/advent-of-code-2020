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
		{name: "0,3,6", startingNumbers: []int{0, 3, 6}, terminationRound: 2020, expectedResult: 436},
		{name: "1,3,2", startingNumbers: []int{1, 3, 2}, terminationRound: 2020, expectedResult: 1},
		{name: "2,1,3", startingNumbers: []int{2, 1, 3}, terminationRound: 2020, expectedResult: 10},
		{name: "1,2,3", startingNumbers: []int{1, 2, 3}, terminationRound: 2020, expectedResult: 27},
		{name: "2,3,1", startingNumbers: []int{2, 3, 1}, terminationRound: 2020, expectedResult: 78},
		{name: "3,2,1", startingNumbers: []int{3, 2, 1}, terminationRound: 2020, expectedResult: 438},
		{name: "3,1,2", startingNumbers: []int{3, 1, 2}, terminationRound: 2020, expectedResult: 1836},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedResult, partOne(test.startingNumbers, test.terminationRound))
		})
	}
}
