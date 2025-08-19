package days

import (
	u "aoc2024-go/utils"
)

func Day4() u.Answers {
	contents := u.ReadFileToLines("resources/day4.txt")
	gridPoints := u.PointsInGrid(len(contents[0]), len(contents))
	part1 := part1(gridPoints, contents)
	part2 := part2(gridPoints, contents)
	return u.IntAnswers(part1, part2)
}

func ScaledCompassPoints(scale int) [][]u.Point {
	points := u.AllCompassPoints()
	scaledPoints := make([][]u.Point, len(points))
	for i, point := range points {
		scaledPoints[i] = make([]u.Point, scale)
		for j := range scale {
			scaledPoints[i][j] = u.ScalePoint(point, j)
		}
	}
	return scaledPoints
}

func part1(gridPoints []u.Point, contents []string) int {
	scaledCompassPoints := ScaledCompassPoints(len("XMAS"))
	return runPart(gridPoints, contents, scaledCompassPoints, []string{"XMAS"})
}

func part2(gridPoints []u.Point, contents []string) int {
	xmases := []string{"SSAMM", "MSAMS", "SMASM", "MMASS"}
	xDirections := []u.Point{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: 0, Y: 0},
		{X: -1, Y: 1},
		{X: 1, Y: 1},
	}
	return runPart(gridPoints, contents, [][]u.Point{xDirections}, xmases)
}

func runPart(gridPoints []u.Point, contents []string, lines [][]u.Point, matches []string) int {
	count := 0
	for p := range gridPoints {
		pos := gridPoints[p]
		for _, line := range lines {
			count += countMatches(contents, pos, line, matches)
		}
	}
	return count
}

func countMatches(lines []string, pos u.Point, offsets []u.Point, words []string) int {
	gridWord := string(gather(lines, pos, offsets))
	for _, word := range words {
		if gridWord == word {
			return 1
		}
	}
	return 0
}

func gather(lines []string, pos u.Point, diffs []u.Point) []byte {
	wordLen := len(diffs)
	width := len(lines[0])
	result := make([]byte, wordLen)
	for i, diff := range diffs {
		x := pos.X + diff.X
		y := pos.Y + diff.Y
		if x < 0 || x >= width || y < 0 || y >= len(lines) {
			return nil
		}
		result[i] = lines[y][x]
	}
	return result
}
