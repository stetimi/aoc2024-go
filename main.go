package main

import (
	"aoc2024-go/days"
	"aoc2024-go/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dayFuncs := []utils.DayFunc{
		days.Day1,
		days.Day2,
		days.Day3,
		days.Day4,
	}
	var selectedDays []int
	if len(os.Args) == 1 {
		selectedDays = rangeArrayFrom1(len(dayFuncs))
	} else {
		selectedDays = parseDayArgs(os.Args[1], len(dayFuncs))
	}
	for _, day := range selectedDays {
		start := time.Now()
		answers := dayFuncs[day-1]()
		elapsed := time.Since(start).Seconds() * 1000
		fmt.Printf("Day %d [%.2f ms]: ", day, elapsed)
		fmt.Printf("Part 1: %s; Part 2: %s\n", answers.Part1, answers.Part2)
	}
}

func rangeArrayFrom1(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}

func parseDayArgs(arg string, maxDay int) []int {
	daysList := []int{}
	for _, part := range strings.Split(arg, ",") {
		trimmed := strings.TrimSpace(part)
		day, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(fmt.Sprintf("Invalid integer: %s", trimmed))
		}
		if day < 1 || day > maxDay {
			panic(fmt.Sprintf("Day out of range: %d. Please provide integers between 1 and %d", day, maxDay))
		}
		daysList = append(daysList, day)
	}
	if len(daysList) == 0 {
		panic("No valid day numbers provided.")
	}
	return daysList
}
