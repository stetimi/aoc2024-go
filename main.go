package main

import (
	"aoc2024-go/days"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dayFuncs := []func(){
		days.Day1,
		days.Day2,
		days.Day3,
		days.Day4,
	}
	if len(os.Args) == 1 {
		fmt.Println("No arguments provided. Running all days...")
		for i, f := range dayFuncs {
			start := time.Now()
			f()
			elapsed := time.Since(start).Seconds() * 1000
			fmt.Printf("Day %d [%.2f ms]\n", i+1, elapsed)
		}
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <day_number>[,<day_number>...]")
		os.Exit(1)
	}
	daysList := parseDayArgs(os.Args[1], len(dayFuncs))
	for _, day := range daysList {
		start := time.Now()
		dayFuncs[day-1]()
		elapsed := time.Since(start).Seconds() * 1000
		fmt.Printf("Day %d [%.2f ms]\n", day, elapsed)
	}
}

func parseDayArgs(arg string, maxDay int) []int {
	daysList := []int{}
	for _, part := range strings.Split(arg, ",") {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
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
