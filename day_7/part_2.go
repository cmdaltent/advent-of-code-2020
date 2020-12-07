package main

import (
	"log"
	"os"
)

func partTwo(color string, input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}
	bagRules, err := parseBagRules(file)
	if err != nil {
		log.Fatalln(err)
	}

	return countChildren(color, bagRules) - 1
}

func countChildren(color string, rules bagRules) int {
	counter := 1
	for color, count := range rules[color] {
		counter += count * countChildren(color, rules)
	}
	return counter
}
