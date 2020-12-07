package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partOne(t *testing.T) {
	assert.Equal(t, 4, partOne("shiny gold", "testdata/part_1_example.txt"))
}
