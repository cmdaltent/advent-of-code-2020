package main

import "github.com/adam-lavrik/go-imath/ix"

func partTWo(input []int, weakness int) int {
	for o, outer := range input {
		for i := o + 1; i < len(input); i++ {
			outer += input[i]
			if outer == weakness {
				return ix.MinSlice(input[o:i+1]) + ix.MaxSlice(input[o:i+1])
			}
		}
	}
	return 0
}
