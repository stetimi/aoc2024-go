package days

import (
	"aoc2024-go/utils"
	"strconv"
	"strings"
)

func Day2() utils.Answers {
	contents := utils.ReadFileToLines("resources/day2.txt")
	reports := make([][]int, len(contents))
	for i, report_text := range contents {
		reports[i] = splitToInts(report_text)
	}

	safe_reports_day1 := 0
	safe_reports_day2 := 0
	for _, report := range reports {
		if isSafe(report, -1) {
			safe_reports_day1++
			safe_reports_day2++
			continue
		}
		for ignore_index := 0; ignore_index < len(report); ignore_index++ {
			if isSafe(report, ignore_index) {
				safe_reports_day2++
				break
			}
		}
	}
	return utils.IntAnswers(safe_reports_day1, safe_reports_day2)
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

func isSafe(report []int, ignore_index int) bool {
	working_report := report
	if ignore_index != -1 {
		working_report = ignoreIndex(report, ignore_index)
	}
	direction := utils.Sign(working_report[1] - working_report[0])
	if direction == 0 {
		return false
	}
	for i := 0; i < len(working_report)-1; i++ {
		diff := working_report[i+1] - working_report[i]
		if utils.Sign(diff) != direction {
			return false
		}
		abs_diff := utils.Abs(diff)
		if abs_diff > 3 {
			return false
		}
	}
	return true
}

func ignoreIndex(report []int, ignoredIndex int) []int {
	copy := make([]int, len(report)-1)
	ci := 0
	for ri := 0; ri < len(report); ri++ {
		if ri != ignoredIndex {
			copy[ci] = report[ri]
			ci++
		}
	}
	return copy
}
