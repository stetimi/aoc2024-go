package days

import (
	"aoc2024-go/utils"
	"strconv"
	"strings"
)

func Day2() {
	reports := utils.ReadFileToLines("resources/day2.txt")
	safe_reports := 0
	for _, report_text := range reports {
		report := splitToInts(report_text)
		if isSafe(report) {
			safe_reports++
		}
	}
	println("Part 1: ", safe_reports)
}

func splitToInts(report string) []int {
	level_strs := strings.Split(report, " ")
	levels := make([]int, len(level_strs))
	for i, level_str := range level_strs {
		level, err := strconv.Atoi(level_str)
		if err != nil {
			panic("Invalid level in report: " + level_str)
		}
		levels[i] = level
	}
	return levels
}

func isSafe(report []int) bool {
	requiredSign := utils.Sign(report[1] - report[0])
	if requiredSign == 0 {
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if utils.Sign(diff) != requiredSign || utils.Abs(diff) > 3 {
			return false
		}
	}
	return true
}
