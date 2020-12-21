package part_2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartTwo(filePath string) uint64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	addressSpace := make(map[string]uint64)
	var bitMask string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			splits := strings.Split(line, " = ")
			bitMask = splits[1]
			continue
		}
		if strings.HasPrefix(line, "mem") {
			binaryMemAddr, binary, err := parseValueLine(line)
			if err != nil {
				panic(err)
			}
			for _, bma := range merge(binaryMemAddr, bitMask) {
				addressSpace[bma] = binary
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var sum uint64
	for _, v := range addressSpace {
		sum += v
	}
	return sum
}

func parseValueLine(line string) (binaryMemAddr string, value uint64, err error) {
	binaryMemAddr, err = parseMemoryAddressAsBinary(line)
	if err != nil {
		return
	}
	value, err = parseValue(line)
	return
}

func parseMemoryAddressAsBinary(line string) (string, error) {
	regex := *regexp.MustCompile(`\[(\d*)]`)
	match := regex.FindString(line)
	match = strings.ReplaceAll(match, "[", "")
	match = strings.ReplaceAll(match, "]", "")
	parsed, err := strconv.ParseUint(match, 10, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(parsed, 2), nil
}

func parseValue(line string) (uint64, error) {
	index := strings.Index(line, "=") + 2
	return strconv.ParseUint(line[index:], 10, 64)
}

func merge(binMemAddr string, bitmask string) []string {
	addrLen := len(binMemAddr)
	maskLen := len(bitmask)

	merged := bitmask[:maskLen-addrLen]
	bitmask = bitmask[maskLen-addrLen:]
	for i := 0; i < addrLen; i++ {
		switch bitmask[i] {
		case '0':
			merged += string(binMemAddr[i])
		case '1', 'X':
			merged += string(bitmask[i])
		}
	}
	return replaceX(merged)
}

func replaceX(floatingAddr string) []string {
	var addresses []string
	if index := strings.Index(floatingAddr, "X"); index > -1 {
		with0 := floatingAddr[:index] + "0" + floatingAddr[index+1:]
		with1 := floatingAddr[:index] + "1" + floatingAddr[index+1:]
		addresses = append(addresses, replaceX(with0)...)
		addresses = append(addresses, replaceX(with1)...)
	} else {
		addresses = append(addresses, floatingAddr)
	}
	return addresses
}
