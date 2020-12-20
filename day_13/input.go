package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readSchedule(filePath string) (earliestDeparture int, busLines []int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read earliest departure, which is the first line of the input
	if scanner.Scan() {
		earliestDeparture, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
	}

	// read the bus lines, which is the second line of the input
	if scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		for _, bus := range lines {
			if bus != "x" {
				busLine, err := strconv.Atoi(bus)
				if err != nil {
					panic(err)
				}
				busLines = append(busLines, busLine)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
