package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseBagRules(file *os.File) (bagRules, error) {
	bagRules := make(bagRules)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bg := parseBagRule(scanner.Text())
		bagRules[bg.color] = bg.children
	}
	return bagRules, scanner.Err()
}

func parseBagRule(rule string) bagRule {
	elements := strings.Split(rule, " ")
	if strings.HasSuffix(rule, "no other bags.") {
		return bagRule{
			color: strings.Join(elements[:2], " "),
		}
	}

	children := make(map[string]int)

	idx := 4
	for idx < len(elements) {
		amount, err := strconv.Atoi(elements[idx])
		if err != nil {
			log.Fatalln(err)
		}
		children[strings.Join(elements[idx+1:idx+3], " ")] = amount
		idx += 4
	}
	return bagRule{
		color:    strings.Join(elements[:2], " "),
		children: children,
	}
}

type bagRules map[string]map[string]int

type bagRule struct {
	color    string
	children map[string]int
}
