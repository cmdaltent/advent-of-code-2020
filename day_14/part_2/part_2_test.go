package part_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replaceX(t *testing.T) {
	assert.Len(t, replaceX("0111X10100100X1111X10010X000X1000001"), 32)
}
