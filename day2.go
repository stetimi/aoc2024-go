package main

import (
	"aoc2024-go/utils"
)

func Main2() {
	reports := utils.ReadFileToLines("resources/day2.txt")
	println("Part 1: ", reports[1])
}

func isSafe(report []int) bool {
	requiredSign := utils.Sign(report[1] - report[0])
	if requiredSign == 0 {
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if utils.Sign(diff) != requiredSign || utils.Abs(diff) > 2 {
			return false
		}
	}
	return true
}
