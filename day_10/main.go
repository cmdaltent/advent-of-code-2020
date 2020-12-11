package main

import (
	"log"

	"github.com/cmdaltent/advent-of-code-2020/lib"
)

func main() {
	input := lib.ReadIntSliceFromFile("day_10/input.txt")

	log.Printf("Part One: Multiplication Value: %d\n", partOne(input))
}
