package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func partTwo() {
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

		first, second, found := checkThreeFor2020(numbers, current)
		if found {
			log.Printf("Solution: %d + %d + %d = 2020 and %d * %d * %d = %d\n", current, first, second, current, first, second, current*first*second)
			break
		} else {
			numbers = append(numbers, current)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkThreeFor2020(n []int64, c int64) (int64, int64, bool) {
	l := len(n)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if n[i]+n[j]+c == 2020 {
				return n[i], n[j], true
			}
		}
	}

	return 0, 0, false
}
