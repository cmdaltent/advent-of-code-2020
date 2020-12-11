package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_9/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		input = append(input, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	weakness := partOne(input, 25)
	log.Printf("Part One: Weakness: %d\n", weakness)
	log.Printf("Part Two: Sum: %d\n", partTWo(input, weakness))
}
