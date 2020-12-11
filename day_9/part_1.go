package main

func partOne(input []int, preamble int) int {
	weakness := -1
	counter := preamble

	for counter < len(input) {
		number := input[counter]
		lower := counter - preamble

	outer:
		for lower < counter-1 {
			for i := lower + 1; i < counter; i++ {
				if input[lower]+input[i] == number {
					break outer
				}
			}
			lower++
		}

		if lower == counter-1 {
			weakness = number
			break
		}

		counter++
	}

	return weakness
}
