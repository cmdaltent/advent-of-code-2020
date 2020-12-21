package part_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseMemoryAddress(t *testing.T) {
	assert.Equal(t, "44861", parseMemoryAddress("mem[44861] = 522064487"))
}

func Test_parseValueAsBinary(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectError  bool
		expectedMask string
	}{
		{name: "All Good, expect 100", input: "mem[15362] = 4", expectError: false, expectedMask: "100"},
		{name: "Invalid Value, error", input: "mem[15362] = x34", expectError: true, expectedMask: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			binary, err := parseValueAsBinary(test.input)
			assert.Equal(t, test.expectedMask, binary)
			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	assert.Equal(t, "011101010010001111010010100001000001", merge("100001000001", "0111X10100100X1111X10010X000X1000001"))
}
