package main

import (
	"log"
	"os"
)

func partOne(color string, input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	bagRules, err := parseBagRules(file)
	if err != nil {
		log.Fatalln(err)
	}

	var count int
	for ruleColor := range bagRules {
		if canContain(color, ruleColor, bagRules) {
			count++
		}
	}
	return count
}

func canContain(searched, current string, rules bagRules) bool {
	var found bool
	for color := range rules[current] {
		if color == searched {
			return true
		}
		found = found || canContain(searched, color, rules)
	}
	return found
}
