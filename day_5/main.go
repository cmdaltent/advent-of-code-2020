package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var highestSeatID int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		rowDef := boardingPass[:7]
		colDef := boardingPass[7:]

		var wg sync.WaitGroup
		wg.Add(2)

		var (
			row    int
			col    int
			rowErr error
			colErr error
		)

		go func() {
			rows := generateRows()
			row, rowErr = determineRow(rows, rowDef)
			wg.Done()
		}()

		go func() {
			cols := generateCols()
			col, colErr = determineCol(cols, colDef)
			wg.Done()
		}()

		wg.Wait()
		if rowErr != nil || colErr != nil {
			log.Fatalf("RowErr: %s\nColErr: %s\n", rowErr, colErr)
		}

		seatID := (row * 8) + col
		if highestSeatID < seatID {
			highestSeatID = seatID
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Highest seat id: %d\n", highestSeatID)
}

func generateRows() []int {
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}
	return rows
}

func generateCols() []int {
	cols := make([]int, 8)
	for i := 0; i < 8; i++ {
		cols[i] = i
	}
	return cols
}

var errInvalidBoardingPass = fmt.Errorf("invalid boarding pass")

func determineRow(rows []int, boardingPass string) (int, error) {
	if len(rows) == 1 {
		return rows[0], nil
	}

	half := len(rows) / 2

	switch boardingPass[0] {
	case 'F':
		return determineRow(rows[:half], boardingPass[1:])
	case 'B':
		return determineRow(rows[half:], boardingPass[1:])
	}

	return 0, errInvalidBoardingPass
}

func determineCol(cols []int, boardingPass string) (int, error) {
	if len(cols) == 1 {
		return cols[0], nil
	}

	half := len(cols) / 2

	switch boardingPass[0] {
	case 'L':
		return determineCol(cols[:half], boardingPass[1:])
	case 'R':
		return determineCol(cols[half:], boardingPass[1:])
	}

	return 0, errInvalidBoardingPass
}
