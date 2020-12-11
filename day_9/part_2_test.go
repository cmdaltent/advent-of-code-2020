package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partTwo(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	weakness := 127
	assert.Equal(t, 62, partTWo(input, weakness))
}
