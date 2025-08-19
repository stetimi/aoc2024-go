package days

import (
	u "aoc2024-go/utils"
)

func Day4() u.Answers {
	contents := u.ReadFileToLines("resources/day4.txt")
	gridPoints := u.PointsInGrid(len(contents[0]), len(contents))
	count := 0
	for p := range gridPoints {
		pos := gridPoints[p]
		count += countMatches(contents, pos, "XMAS")
	}
	return u.IntAnswers(count, -1)
}

func countMatches(lines []string, pos u.Point, word string) int {
	count := 0
	for _, direction := range u.AllCompassPoints() {
		if hasMatchAlongDirection(lines, pos, direction, word) {
			count++
		}
	}
	return count
}

func hasMatchAlongDirection(lines []string, pos u.Point, direction u.Point, word string) bool {
	word_len := len(word)
	width := len(lines[0])
	last_x := pos.X + direction.X*(word_len-1)
	if last_x < 0 || last_x >= width {
		return false
	}
	last_y := pos.Y + direction.Y*(word_len-1)
	if last_y < 0 || last_y >= len(lines) {
		return false
	}
	for i := range word_len {
		x := pos.X + direction.X*i
		y := pos.Y + direction.Y*i
		if lines[y][x] != word[i] {
			return false
		}
	}
	return true
}
