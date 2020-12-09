package main

type operation uint8

const (
	acc operation = iota
	jmp
	nop
)

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
