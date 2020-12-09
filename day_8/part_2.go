package main

func partTwo(instructionSet []instruction) int {
	var accumulator int
	for idx := range instructionSet {
		inst := instructionSet[idx]
		if inst.operation == acc {
			continue
		}

		cpy := make([]instruction, len(instructionSet))
		copy(cpy, instructionSet)

		cpy[idx].operation = inst.operation.invert()

		pp := programPointer{
			instructionSet: cpy,
		}
		for pp.next() {
		}

		if pp.currentIndex == len(instructionSet) {
			accumulator = pp.accumulator
			break
		}
	}
	return accumulator
}

func reset(input []instruction) []instruction {
	cpy := make([]instruction, len(input))
	for i, e := range input {
		cpy[i] = instruction{
			operation: e.operation,
			value:     e.value,
		}
	}
	return cpy
}
