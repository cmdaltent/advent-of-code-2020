package part_1

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/cmdaltent/advent-of-code-2020/day_5/common"
)

func PartOne() {
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
			rows := common.GenerateRows()
			row, rowErr = common.DetermineRow(rows, rowDef)
			wg.Done()
		}()

		go func() {
			cols := common.GenerateCols()
			col, colErr = common.DetermineCol(cols, colDef)
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
