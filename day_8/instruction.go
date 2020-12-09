package main

type operation uint8

const (
	acc operation = iota
	jmp
	nop
)

func (o operation) invert() operation {
	switch o {
	default:
		return o
	case jmp:
		return nop
	case nop:
		return jmp
	}
}

type instruction struct {
	operation          operation
	value              int
	numberOfExecutions int
}

type programPointer struct {
	accumulator    int
	currentIndex   int
	instructionSet []instruction
}

func (pp *programPointer) next() bool {
	if pp.currentIndex >= len(pp.instructionSet) {
		return false
	}

	inst := pp.instructionSet[pp.currentIndex]
	if inst.numberOfExecutions > 0 {
		return false
	}
	pp.instructionSet[pp.currentIndex].numberOfExecutions += 1

	switch inst.operation {
	case acc:
		pp.currentIndex += 1
		pp.accumulator += inst.value
	case nop:
		pp.currentIndex += 1
	case jmp:
		pp.currentIndex += inst.value
	}
	return true
}
