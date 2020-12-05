package main

import "log"

func main() {
	log.Printf("Part one: Encountered trees: %d\n", traverse(shift{1, 3}))
	log.Printf("Part two: Multiplied number of encountered trees: %d\n", partTwo())
}
