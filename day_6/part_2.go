package main

import (
	"bufio"
	"log"
	"os"
)

func partTwo() {
	file, err := os.Open("day_6/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var answerCount int
	group := newEveryoneGroupAnswers()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" {
			answerCount += group.groupCount()
			group = newEveryoneGroupAnswers()
			continue
		}
		group.incGroupMembers()
		for i := 0; i < len(line); i++ {
			group.registerAnswer(line[i])
		}
	}
	answerCount += group.groupCount()

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Sum of Answers: %d\n", answerCount)
}

func newEveryoneGroupAnswers() everyoneGroupAnswers {
	return everyoneGroupAnswers{
		answers: map[uint8]int{},
	}
}

type everyoneGroupAnswers struct {
	numberOfGroupMembers int
	answers              map[uint8]int
}

func (ega *everyoneGroupAnswers) incGroupMembers() {
	ega.numberOfGroupMembers = ega.numberOfGroupMembers + 1
}

func (ega *everyoneGroupAnswers) registerAnswer(answer uint8) {
	if count, ok := ega.answers[answer]; ok {
		ega.answers[answer] = count + 1
	} else {
		ega.answers[answer] = 1
	}
}

func (ega *everyoneGroupAnswers) groupCount() int {
	var count int
	for _, answeredBy := range ega.answers {
		if answeredBy == ega.numberOfGroupMembers {
			count++
		}
	}
	return count
}
