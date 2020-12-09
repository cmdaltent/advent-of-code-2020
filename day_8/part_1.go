package main

func partOne(instructionSet []instruction) int {
	pp := programPointer{
		instructionSet: instructionSet,
	}

	for pp.next() {
	}
	return pp.accumulator
}
