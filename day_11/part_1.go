package main

type adjacent struct {
	row    int
	column int
}

func partOne(seatMap [][]seatStatus) int {

	var isStable bool
	var occupiedSeats int

	for !isStable {
		nextGen := make([][]seatStatus, len(seatMap))
		for row, seats := range seatMap {
			nextGen[row] = make([]seatStatus, len(seats))
			for column, seat := range seats {
				// 2 3 4
				// 1 * 5
				// 8 7 6
				adjacents := []adjacent{
					{row: row, column: column - 1},     // 1
					{row: row - 1, column: column - 1}, // 2
					{row: row - 1, column: column},     // 3
					{row: row - 1, column: column + 1}, // 4
					{row: row, column: column + 1},     // 5
					{row: row + 1, column: column + 1}, // 6
					{row: row + 1, column: column},     // 7
					{row: row + 1, column: column - 1}, // 8
				}
				switch seat {
				case empty:
					nextGen[row][column] = newSeatStatusForEmpty(adjacents, seatMap)
				case occupied:
					nextGen[row][column] = newSeatStatusForOccupied(adjacents, seatMap)
				case floor:
					nextGen[row][column] = floor
				}
			}
		}
		if os := countOccupiedSeats(nextGen); os != occupiedSeats {
			occupiedSeats = os
			seatMap = nextGen
		} else {
			isStable = true
		}
	}

	return occupiedSeats
}

func newSeatStatusForEmpty(adjacents []adjacent, seatMap [][]seatStatus) seatStatus {
	for _, a := range adjacents {
		if a.row < 0 || a.row >= len(seatMap) {
			continue
		}
		if a.column < 0 || a.column >= len(seatMap[a.row]) {
			continue
		}
		if seatMap[a.row][a.column] == occupied {
			return empty
		}
	}
	return occupied
}

func newSeatStatusForOccupied(adjacents []adjacent, seatMap [][]seatStatus) seatStatus {
	var occupieds int
	for _, a := range adjacents {
		if a.row < 0 || a.row >= len(seatMap) {
			continue
		}
		if a.column < 0 || a.column >= len(seatMap[a.row]) {
			continue
		}
		if seatMap[a.row][a.column] == occupied {
			occupieds++
		}
	}
	if occupieds >= 4 {
		return empty
	}
	return occupied
}

func countOccupiedSeats(seatMap [][]seatStatus) int {
	var occupiedSeats int
	for _, seats := range seatMap {
		for _, seat := range seats {
			if seat == occupied {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}
