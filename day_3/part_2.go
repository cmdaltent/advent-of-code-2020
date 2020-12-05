package main

func partTwo() int {
	shifts := []shift{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	encounteredTrees := 1
	for _, s := range shifts {
		encounteredTrees *= traverse(s)
	}
	return encounteredTrees
}
