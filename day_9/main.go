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

	log.Printf("Part One: Weakness: %d\n", partOne(input, 25))
}
