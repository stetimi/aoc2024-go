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
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day number. Please provide an integer.")
		os.Exit(1)
	}
	switch day {
	case 1:
		days.Day1()
	case 2:
		days.Day2()
	default:
		fmt.Printf("Day %d not implemented.\n", day)
		os.Exit(1)
	}
}
