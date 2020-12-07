package main

import "log"

func main() {
	log.Println("Part one")
	partOneBagColor := "shiny gold"
	log.Printf("# of Bags in which a '%s' bag fits: %d.\n", partOneBagColor, partOne(partOneBagColor, "day_7/input.txt"))

	log.Println("Partn two")
	log.Printf("A '%s' carries %d bags.\n", partOneBagColor, partTwo(partOneBagColor, "day_7/input.txt"))
}
