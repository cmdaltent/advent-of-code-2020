package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func partOne() {
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var passports []passport

	err = readPassports(file, func(pp passport) {
		passports = append(passports, pp)
	})
	if err != nil {
		log.Fatalln(err)
	}

	var validPassports int
	for i := range passports {
		if passports[i].isValid() {
			validPassports++
		}
	}

	log.Printf("Read Passports: %d\n", len(passports))
	log.Printf("Valid Passports: %d\n", validPassports)
}

type passportParsed func(pp passport)

func readPassports(file *os.File, cb passportParsed) error {
	var pp passport

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line == "\n" {
			cb(pp)
			pp = passport{}
			continue
		}
		for _, elem := range strings.Split(line, " ") {
			kv := strings.Split(elem, ":")
			pp.mergeWith(kv[0], kv[1])
		}
	}
	cb(pp)
	return scanner.Err()
}
