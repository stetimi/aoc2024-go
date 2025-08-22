package main

import (
	"aoc2024-go/days"
	"aoc2024-go/utils"
	_ "embed"
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
		days.Day5,
		days.Day6,
		days.Day7,
		days.Day8,
	}
	inputs := inputs()
	var selectedDays []int
	if len(os.Args) == 1 {
		selectedDays = rangeArrayFrom1(len(dayFuncs))
	} else {
		selectedDays = parseDayArgs(os.Args[1], len(dayFuncs))
	}
	total_elapsed := 0.0
	for _, day := range selectedDays {
		start := time.Now()
		answers := dayFuncs[day-1](inputs[day-1])
		elapsed := time.Since(start).Seconds() * 1000
		total_elapsed += elapsed
		fmt.Printf("Day %d [%.2f ms]: ", day, elapsed)
		fmt.Printf("Part 1: %s; Part 2: %s\n", answers.Part1, answers.Part2)
	}
	fmt.Printf("Total elapsed time: %.2f ms\n", total_elapsed)
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

//go:embed resources/day1.txt
var day1_file []byte

//go:embed resources/day2.txt
var day2_file []byte

//go:embed resources/day3.txt
var day3_file []byte

//go:embed resources/day4.txt
var day4_file []byte

//go:embed resources/day5.txt
var day5_file []byte

//go:embed resources/day6.txt
var day6_file []byte

//go:embed resources/day7.txt
var day7_file []byte

//go:embed resources/day8.txt
var day8_file []byte

func inputs() [][]byte {
	return [][]byte{
		day1_file,
		day2_file,
		day3_file,
		day4_file,
		day5_file,
		day6_file,
		day7_file,
		day8_file,
	}
}
