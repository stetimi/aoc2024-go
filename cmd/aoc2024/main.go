package main

import (
	"aoc2024-go/internal/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day1, err := os.ReadFile("resources/day1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(day1), "\n")

	left, right := splitIntoTwoColumns(lines, "   ")
	sort.Ints(left)
	sort.Ints(right)
	part1 := sumDiffs(left, right)
	println("Part 1: ", part1)

	rightFrequencies := utils.FrequencyMap(right)
	part2 := calculateSimilarityScore(left, rightFrequencies)
	println("Part 2: ", part2)
}

func splitIntoTwoColumns(lines []string, sep string) ([]int, []int) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		cols := strings.SplitN(line, "   ", 2)
		left[i], _ = strconv.Atoi(cols[0])
		right[i], _ = strconv.Atoi(cols[1])
	}
	return left, right
}

func sumDiffs(left, right []int) int {
	sum := 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// ...existing code...

func calculateSimilarityScore(values []int, frequencies map[int]int) int {
	score := 0
	for _, value := range values {
		score += value * frequencies[value]
	}
	return score
}
