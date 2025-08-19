package days

import (
	"aoc2024-go/utils"
	"sort"
	"strconv"
	"strings"
)

func Day1() utils.Answers {
	lines := utils.ReadFileToLines("resources/day1.txt")
	left, right := splitIntoTwoColumns(lines, "   ")
	sort.Ints(left)
	sort.Ints(right)
	part1 := sumDiffs(left, right)
	rightFrequencies := utils.FrequencyMap(right)
	part2 := calculateSimilarityScore(left, rightFrequencies)
	return utils.IntAnswers(part1, part2)
}

func splitIntoTwoColumns(lines []string, sep string) ([]int, []int) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		cols := strings.SplitN(line, sep, 2)
		left[i], _ = strconv.Atoi(cols[0])
		right[i], _ = strconv.Atoi(cols[1])
	}
	return left, right
}

func sumDiffs(left, right []int) int {
	sum := 0
	for i := range left {
		sum += utils.Abs(left[i] - right[i])
	}
	return sum
}

func calculateSimilarityScore(values []int, frequencies map[int]int) int {
	score := 0
	for _, value := range values {
		score += value * frequencies[value]
	}
	return score
}
