package common

import "fmt"

func GenerateRows() []int {
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}
	return rows
}

func GenerateCols() []int {
	cols := make([]int, 8)
	for i := 0; i < 8; i++ {
		cols[i] = i
	}
	return cols
}

var ErrInvalidBoardingPass = fmt.Errorf("invalid boarding pass")

func DetermineRow(rows []int, boardingPass string) (int, error) {
	if len(rows) == 1 {
		return rows[0], nil
	}

	half := len(rows) / 2

	switch boardingPass[0] {
	case 'F':
		return DetermineRow(rows[:half], boardingPass[1:])
	case 'B':
		return DetermineRow(rows[half:], boardingPass[1:])
	}

	return 0, ErrInvalidBoardingPass
}

func DetermineCol(cols []int, boardingPass string) (int, error) {
	if len(cols) == 1 {
		return cols[0], nil
	}

	half := len(cols) / 2

	switch boardingPass[0] {
	case 'L':
		return DetermineCol(cols[:half], boardingPass[1:])
	case 'R':
		return DetermineCol(cols[half:], boardingPass[1:])
	}

	return 0, ErrInvalidBoardingPass
}
