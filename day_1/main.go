package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		summand, found := checkFor2020(numbers, current)
		if found {
			log.Printf("Solution: %d + %d = 2020 and %d * %d = %d\n", current, summand, current, summand, current*summand)
			break
		} else {
			numbers = append(numbers, current)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkFor2020(n []int64, c int64) (int64, bool) {
	for i := range n {
		if n[i]+c == 2020 {
			return n[i], true
		}
	}
	return 0, false
}
