package main

import "log"

func main() {
	log.Printf("Part One - 2020th number: %d\n", play([]int{0, 1, 5, 10, 3, 12, 19}, 2020))
	log.Printf("Part One - 30000000th number: %d\n", play([]int{0, 1, 5, 10, 3, 12, 19}, 30000000))
}
