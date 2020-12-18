package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	navigationInstructions, err := readNavigationInstructions("testdata/partOne_test_input.txt")
	require.NoError(t, err)

	assert.Equal(t, 25, partOne(navigationInstructions))
}
