package utils

import (
	"strconv"
)

type Answers struct {
	Part1 string
	Part2 string
}

type DayFunc func([]byte) Answers

func IntAnswers(part1, part2 int) Answers {
	return Answers{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func StringAnswers(part1, part2 string) Answers {
	return Answers{
		Part1: part1,
		Part2: part2,
	}
}

func TodoAnswers() Answers {
	return Answers{
		Part1: "TODO",
		Part2: "TODO",
	}
}
