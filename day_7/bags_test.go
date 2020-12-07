package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseBagRule(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bagRule
	}{
		{"No child bags", "faded blue bags contain no other bags.", bagRule{color: "faded blue"}},
		{"One child bag", "bright white bags contain 1 shiny gold bag.", bagRule{color: "bright white", children: map[string]int{"shiny gold": 1}}},
		{"Many children bags", "dark olive bags contain 3 faded blue bags, 4 dotted black bags.", bagRule{color: "dark olive", children: map[string]int{"faded blue": 3, "dotted black": 4}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, parseBagRule(test.input))
		})
	}
}
