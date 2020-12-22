package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/cmdaltent/advent-of-code-2020/day_16/common"
)

func ReadInput(filePath string) (fields map[string]common.Constraint, myTicket []int, nearbyTickets [][]int, err error) {
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

func parseFields(scanner *bufio.Scanner) (map[string]common.Constraint, error) {
	fields := make(map[string]common.Constraint)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		splits := strings.Split(line, ": ")
		intervals := strings.Split(splits[1], " or ")

		minMax1 := strings.Split(intervals[0], "-")
		min1, err := strconv.Atoi(minMax1[0])
		if err != nil {
			return nil, err
		}
		max1, err := strconv.Atoi(minMax1[1])
		if err != nil {
			return nil, err
		}

		minMax2 := strings.Split(intervals[1], "-")
		min2, err := strconv.Atoi(minMax2[0])
		if err != nil {
			return nil, err
		}
		max2, err := strconv.Atoi(minMax2[1])
		if err != nil {
			return nil, err
		}

		fields[splits[0]] = common.Constraint{Inter1: common.Interval{Min: min1, Max: max1}, Inter2: common.Interval{Min: min2, Max: max2}}
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
