package main

import "log"

func main() {
	seatMap, err := readSeatMap("day_11/input.txt")
	if err != nil {
		panic(err)
	}
	log.Printf("Part One - Occupied Seats: %d\n", partOne(seatMap))
}
