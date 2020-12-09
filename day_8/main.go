package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_8/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var instructionSet []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inst, err := parse(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		instructionSet = append(instructionSet, inst)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Accumulator before infinite loop: %d\n", partOne(instructionSet))
}
