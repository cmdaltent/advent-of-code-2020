package main

func play(startingNumbers []int, terminationRound int) int {
	brain := make(map[int][]int)
	for i, number := range startingNumbers {
		brain[number] = append(brain[number], i+1)
	}

	lastNumber := startingNumbers[len(startingNumbers)-1]
	for round := len(startingNumbers) + 1; round <= terminationRound; round++ {
		prev := brain[lastNumber]
		if len(prev) == 1 {
			brain[0] = append(brain[0], round)
			lastNumber = 0
			continue
		}
		lastNumber = prev[len(prev)-1] - prev[len(prev)-2]
		brain[lastNumber] = append(brain[lastNumber], round)
	}

	return lastNumber
}
