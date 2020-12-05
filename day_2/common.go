package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type policyCheck func(string, int, int, string) bool

func checkInput(check policyCheck) {
	file, err := os.Open("day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validPasswords int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		policy, password := parseLine(scanner.Text())
		min, max, char := parsePolicy(policy)
		if check(password, min, max, char) {
			validPasswords++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of valid passwords: %d\n", validPasswords)
}

func parseLine(line string) (policy string, password string) {
	s := strings.Split(line, ": ")
	policy = s[0]
	password = s[1]
	return
}

func parsePolicy(policy string) (int, int, string) {
	s := strings.Split(policy, " ")
	char := s[1]

	o := strings.Split(s[0], "-")

	left, err := strconv.ParseInt(o[0], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	right, err := strconv.ParseInt(o[1], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return int(left), int(right), char
}
