package part_2

import (
	"bufio"
	"log"
	"os"
	"sync"

	"github.com/cmdaltent/advent-of-code-2020/day_5/common"
)

func PartTwo() {
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	seatMap := generateAllSeatsMap()

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

		seatMap[(row*8)+col] = true
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("My seat: %d", findMySeat(seatMap))
}

// generateAllSeatsMap assigns the occupation status of a seat id. `false` means the seat is not taken, i.e., a valid
// boarding pass does not exists, `true` means the seat is occupied.
func generateAllSeatsMap() map[int]bool {
	seatMap := make(map[int]bool, 128*8)

	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
			seatMap[(r*8)+c] = false
		}
	}

	return seatMap
}

func findMySeat(seatMap map[int]bool) int {
	for seatID, occupied := range seatMap {
		if !occupied && seatMap[seatID+1] == true && seatMap[seatID-1] == true {
			return seatID
		}
	}
	return -1
}
