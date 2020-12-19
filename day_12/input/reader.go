package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cmdaltent/advent-of-code-2020/day_12/common"
)

func ReadNavigationInstructions(filePath string) ([]common.NavigationInstruction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var navigationInstructions []common.NavigationInstruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		switch line[:1] {
		case "N":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.N, Value: value})
		case "S":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.S, Value: value})
		case "E":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.E, Value: value})
		case "W":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.W, Value: value})
		case "R":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.R, Value: value})
		case "L":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.L, Value: value})
		case "F":
			navigationInstructions = append(navigationInstructions, common.NavigationInstruction{Action: common.F, Value: value})
		default:
			return nil, fmt.Errorf("invalid Action: %s", line[:1])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return navigationInstructions, nil
}
