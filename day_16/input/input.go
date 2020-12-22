package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/cmdaltent/advent-of-code-2020/day_16/common"
)

func ReadInput(filePath string) (fields map[string][]common.Interval, myTicket []int, nearbyTickets [][]int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fields, err = parseFields(scanner)
	if err != nil {
		return
	}
	for scanner.Scan() {
		switch scanner.Text() {
		case "your ticket:":
			tickets, err := parseTickets(scanner)
			if err != nil {
				return nil, nil, nil, err
			}
			myTicket = tickets[0]
		case "nearby tickets:":
			nearbyTickets, err = parseTickets(scanner)
			if err != nil {
				return
			}
		}
	}

	return
}

func parseFields(scanner *bufio.Scanner) (map[string][]common.Interval, error) {
	fields := make(map[string][]common.Interval)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		splits := strings.Split(line, ": ")
		intervals := strings.Split(splits[1], " or ")
		for _, inter := range intervals {
			minMax := strings.Split(inter, "-")
			min, err := strconv.Atoi(minMax[0])
			if err != nil {
				return nil, err
			}
			max, err := strconv.Atoi(minMax[1])
			if err != nil {
				return nil, err
			}
			fields[splits[0]] = append(fields[splits[0]], common.Interval{Min: min, Max: max})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return fields, nil
}

func parseTickets(scanner *bufio.Scanner) ([][]int, error) {
	var tickets [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var fields []int
		for _, s := range strings.Split(line, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			fields = append(fields, n)
		}
		tickets = append(tickets, fields)
	}
	return tickets, nil
}
