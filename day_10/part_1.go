package main

import "sort"

func partOne(input []int) int {
	sort.Ints(input)

	var (
		oneDiff   int
		threeDiff int
	)

	switch input[0] {
	case 1:
		oneDiff++
	case 3:
		threeDiff++
	}

	for i := 0; i < len(input)-1; i++ {
		switch input[i+1] - input[i] {
		case 1:
			oneDiff++
		case 3:
			threeDiff++
		}
	}

	return oneDiff * (threeDiff + 1)
}
