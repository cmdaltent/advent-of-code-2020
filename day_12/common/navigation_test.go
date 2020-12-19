package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		head     Heading
		degree   int
		expected Heading
	}{
		{name: "North 360 North", head: North, degree: 360, expected: North},
		{name: "South 360 South", head: South, degree: 360, expected: South},
		{name: "North 90 West", head: North, degree: 90, expected: West},
		{name: "West 270 North", head: West, degree: 270, expected: North},
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
		head     Heading
		degree   int
		expected Heading
	}{
		{name: "North 360 North", head: North, degree: 360, expected: North},
		{name: "South 360 South", head: South, degree: 360, expected: South},
		{name: "North 90 East", head: North, degree: 90, expected: East},
		{name: "West 270 South", head: West, degree: 270, expected: South},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.head.TurnRight(test.degree))
		})
	}
}

func TestPosition_RotateCounterClockwise(t *testing.T) {
	tests := []struct {
		name          string
		position      Position
		rotationAngle int
		expected      Position
	}{
		{name: "Simple", position: Position{X: 1, Y: 0}, rotationAngle: 90, expected: Position{X: 0, Y: 1}},
		{name: "270", position: Position{X: 1, Y: 0}, rotationAngle: 270, expected: Position{X: 0, Y: -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.position.RotateCounterClockwise(test.rotationAngle))
		})
	}
}

func TestPosition_RotateClockwise(t *testing.T) {
	tests := []struct {
		name          string
		position      Position
		rotationAngle int
		expected      Position
	}{
		{name: "Simple", position: Position{X: 1, Y: 0}, rotationAngle: 90, expected: Position{X: 0, Y: -1}},
		{name: "Input Example", position: Position{X: 10, Y: 4}, rotationAngle: 90, expected: Position{X: 4, Y: -10}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.position.RotateClockwise(test.rotationAngle))
		})
	}
}
