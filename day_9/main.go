package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/lib"
)

func main() {
	input := lib.ReadIntSliceFromFile("day_9/input.txt")

	weakness := partOne(input, 25)
	log.Printf("Part One: Weakness: %d\n", weakness)
	log.Printf("Part Two: Sum: %d\n", partTWo(input, weakness))
}
