package main

import (
	"bufio"
	"log"
	"os"
)

func partOne() {
	file, err := os.Open("day_6/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var answerCount int
	groupAnswers := make(map[uint8]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" {
			answerCount += len(groupAnswers)
			groupAnswers = make(map[uint8]bool)
			continue
		}
		for i := 0; i < len(line); i++ {
			groupAnswers[line[i]] = true
		}
	}
	answerCount += len(groupAnswers)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Sum of Answers: %d\n", answerCount)
}
