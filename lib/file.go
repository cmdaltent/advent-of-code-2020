package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadIntSliceFromFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		input = append(input, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return input
}
