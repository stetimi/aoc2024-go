package days

import (
	"aoc2024-go/utils"
	"regexp"
	"strconv"
)

const MUL_REGEXP = `mul\(([0-9]+),([0-9]+)\)`

func Day3() {
	program := utils.MustReadFile("resources/day3.txt")
	println("Part 1: ", runDay1Program(program))
	println("Part 2: ", runDay2Program(program))
}

func runDay1Program(memory []byte) int {
	return runProgram(memory, regexp.MustCompile(MUL_REGEXP))
}

func runDay2Program(memory []byte) int {
	return runProgram(memory, regexp.MustCompile(`do\(\)|don't\(\)|`+MUL_REGEXP))
}

func runProgram(memory []byte, re *regexp.Regexp) int {
	matches := re.FindAllSubmatch(memory, -1)
	sum := 0
	enabled := true
	for _, match := range matches {
		switch string(match[0]) {
		case "do()":
			enabled = true
			continue
		case "don't()":
			enabled = false
			continue
		}
		if len(match) != 3 {
			panic("Bad multiplication format")
		}
		if enabled {
			x := atoi(match[1])
			y := atoi(match[2])
			sum += x * y
		}
	}
	return sum
}

func atoi(b []byte) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return n
}
