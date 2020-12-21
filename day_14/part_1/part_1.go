package part_1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(filePath string) uint64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	addressSpace := make(map[string]string)
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
			memoryAddress, binary, err := parseValueLine(line)
			if err != nil {
				panic(err)
			}
			addressSpace[memoryAddress] = merge(binary, bitMask)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var sum uint64
	for _, v := range addressSpace {
		parsed, err := strconv.ParseUint(v, 2, 64)
		if err != nil {
			panic(err)
		}
		sum += parsed
	}
	return sum
}

func parseValueLine(line string) (memoryAddress string, binary string, err error) {
	memoryAddress = parseMemoryAddress(line)
	binary, err = parseValueAsBinary(line)
	return
}

func parseMemoryAddress(line string) string {
	regex := *regexp.MustCompile(`\[(\d*)]`)
	match := regex.FindString(line)
	match = strings.ReplaceAll(match, "[", "")
	return strings.ReplaceAll(match, "]", "")
}

func parseValueAsBinary(line string) (string, error) {
	index := strings.Index(line, "=") + 2
	int64Value, err := strconv.ParseInt(line[index:], 10, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(int64Value, 2), nil
}

func merge(binary string, bitmask string) string {
	binLen := len(binary)
	maskLen := len(bitmask)

	merged := bitmask[:maskLen-binLen]
	merged = strings.ReplaceAll(merged, "X", "0")

	bitmask = bitmask[maskLen-binLen:]
	for i := 0; i < binLen; i++ {
		if bitmask[i] == 'X' {
			merged += string(binary[i])
		} else {
			merged += string(bitmask[i])
		}
	}
	return merged
}
