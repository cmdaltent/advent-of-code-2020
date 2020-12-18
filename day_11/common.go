package main

import (
	"bufio"
	"os"
)

type seatStatus int8

const (
	empty seatStatus = iota
	occupied
	floor
)

func readSeatMap(filePath string) ([][]seatStatus, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var seatMap [][]seatStatus

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var seatRow []seatStatus
		for _, c := range line {
			var ss seatStatus
			switch c {
			case 'L':
				ss = empty
			case '.':
				ss = floor
			case '#':
				ss = occupied
			}
			seatRow = append(seatRow, ss)
		}
		seatMap = append(seatMap, seatRow)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return seatMap, nil
}
