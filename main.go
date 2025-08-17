package main

import (
	"aoc2024-go/days"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <day_number>")
		os.Exit(1)
	}
	dayFuncs := []func(){
		days.Day1,
		days.Day2,
		days.Day3,
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > len(dayFuncs) {
		fmt.Printf("Invalid day number. Please provide an integer between 1 and %d\n", len(dayFuncs))
		os.Exit(1)
	}
	dayFuncs[day-1]()
}
