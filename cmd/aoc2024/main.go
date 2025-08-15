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
	left, right := SplitIntoTwoColumns(lines, "   ")
	sort.Ints(left)
	sort.Ints(right)
	diffs := SumDiffs(left, right)
	println("Result: ", diffs)
}

func SplitIntoTwoColumns(lines []string, sep string) ([]int, []int) {
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		cols := strings.SplitN(line, "   ", 2)
		left[i], _ = strconv.Atoi(cols[0])
		right[i], _ = strconv.Atoi(cols[1])
	}
	return left, right
}

func SumDiffs(left, right []int) int {
	sum := 0
	for i := range left {
		sum += Abs(left[i] - right[i])
	}
	return sum
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
