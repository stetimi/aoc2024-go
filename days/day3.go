package days

import (
	"aoc2024-go/utils"
	"regexp"
	"strconv"
)

func Day3() {
	program := utils.MustReadFile("resources/day3.txt")
	println("Part 1: ", runProgram(program))
}

func runProgram(memory []byte) int {
	re := mulRegexp()
	matches := re.FindAllSubmatch(memory, -1)
	sum := 0
	for _, match := range matches {
		if len(match) != 3 {
			panic("Bad multiplication format")
		}
		x := atoi(match[1])
		y := atoi(match[2])
		sum += x * y
	}
	return sum
}

func mulRegexp() *regexp.Regexp {
	return regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
}

func atoi(b []byte) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return n
}
