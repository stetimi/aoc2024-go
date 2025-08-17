package main

import (
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

	rightFrequencies := frequencyMap(right)
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

func frequencyMap(values []int) map[int]int {
	freqMap := make(map[int]int)
	for _, value := range values {
		freqMap[value]++
	}
	return freqMap
}

func calculateSimilarityScore(values []int, frequencies map[int]int) int {
	score := 0
	for _, value := range values {
		if freq, exists := frequencies[value]; exists {
			score += value * freq
		}
	}
	return score
}
