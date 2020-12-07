package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partTwo(t *testing.T) {
	assert.Equal(t, 126, partTwo("shiny gold", "testdata/part_2_example.txt"))
}
