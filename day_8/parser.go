package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(line string) (instruction, error) {
	elem := strings.Split(line, " ")
	if len(elem) != 2 {
		return instruction{}, fmt.Errorf("invalid number of arguments %v", elem)
	}

	val, err := strconv.Atoi(elem[1])
	if err != nil {
		return instruction{}, err
	}

	var op operation
	switch elem[0] {
	case "acc":
		op = acc
	case "jmp":
		op = jmp
	case "nop":
		op = nop
	default:
		return instruction{}, fmt.Errorf("invalid operation: %s", elem[0])
	}

	return instruction{
		operation:          op,
		value:              val,
		numberOfExecutions: 0,
	}, nil
}
