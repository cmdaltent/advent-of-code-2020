package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/cmdaltent/advent-of-code-2020/day_4/common"
	"github.com/cmdaltent/advent-of-code-2020/day_4/part_1"
)

func main() {
	log.Println("Part One")
	processBatchFile(part_1.GeneratePassport)
}

func processBatchFile(gen common.GeneratePassport) {
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var passports []common.Passport

	err = readPassports(file, gen, func(pp common.Passport) {
		passports = append(passports, pp)
	})
	if err != nil {
		log.Fatalln(err)
	}

	var validPassports int
	for i := range passports {
		if passports[i].IsValid() {
			validPassports++
		}
	}

	log.Printf("Read Passports: %d\n", len(passports))
	log.Printf("Valid Passports: %d\n", validPassports)
}

func readPassports(file *os.File, gp common.GeneratePassport, cb common.PassportParsed) error {
	pp := gp()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line == "\n" {
			cb(pp)
			pp = gp()
			continue
		}
		for _, elem := range strings.Split(line, " ") {
			kv := strings.Split(elem, ":")
			pp.MergeWith(kv[0], kv[1])
		}
	}
	cb(pp)
	return scanner.Err()
}
