package main

import "log"

func main() {
	earliestDeparture, busLines := readSchedule("day_13/input.txt")
	log.Printf("Part One: %d\n", partOne(earliestDeparture, busLines))
}
