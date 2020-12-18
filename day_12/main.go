package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	navigationInstructions, err := readNavigationInstructions("day_12/input.txt")
	if err != nil {
		panic(err)
	}

	log.Printf("Part One - Manhatten Distance: %d", partOne(navigationInstructions))
}

func readNavigationInstructions(filePath string) ([]navigationInstruction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var navigationInstructions []navigationInstruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		switch line[:1] {
		case "N":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: N, value: value})
		case "S":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: S, value: value})
		case "E":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: E, value: value})
		case "W":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: W, value: value})
		case "R":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: R, value: value})
		case "L":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: L, value: value})
		case "F":
			navigationInstructions = append(navigationInstructions, navigationInstruction{act: F, value: value})
		default:
			return nil, fmt.Errorf("invalid action: %s", line[:1])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return navigationInstructions, nil
}
