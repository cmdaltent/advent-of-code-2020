package main

import (
	"bufio"
	"log"
	"os"
)

const tree = '#'

type shift struct {
	row int
	col int
}

func traverse(s shift) int {
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

	if grid == nil {
		log.Fatalln("invalid grid setup")
	}
	colBorder := len(grid[0])

	row := s.row
	col := s.col

	var encounteredTrees int

	for row < len(grid) {
		if grid[row][col] == tree {
			encounteredTrees++
		}

		col = (col + s.col) % colBorder
		row += s.row
	}

	return encounteredTrees
}
