package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	earliestDeparture, lines := readSchedule("testdata/part_1_test_input.txt")
	assert.Equal(t, 295, partOne(earliestDeparture, lines))
}
