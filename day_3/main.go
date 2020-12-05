package main

import (
	"bufio"
	"log"
	"os"
)

const tree = '#'

func main() {
	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var grid [][]byte // rows and cols

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	colBorder := len((grid[0]))

	row := 1
	col := 3

	var encounteredTrees int

	for row < len(grid) {
		if grid[row][col] == tree {
			encounteredTrees++
		}

		col = (col + 3) % colBorder
		row++
	}

	log.Printf("Encountered Trees: %d", encounteredTrees)
}
