package days

import (
	u "aoc2024-go/utils"
	"iter"
	"strings"
)

func Day4(contents []byte) u.Answers {
	lines := strings.Split(string(contents), "\n")
	part1 := day4_part1(lines)
	part2 := day4_part2(lines)
	return u.IntAnswers(part1, part2)
}

func ScaledCompassPoints(scale int) [][]u.Point {
	points := u.AllCompassPoints()
	scaledPoints := make([][]u.Point, len(points))
	for i, point := range points {
		scaledPoints[i] = make([]u.Point, scale)
		for j := range scale {
			scaledPoints[i][j] = point.Scale(j)
		}
	}
	return scaledPoints
}

func day4_part1(contents []string) int {
	gridPoints := u.PointsInGrid(len(contents[0]), len(contents))
	scaledCompassPoints := ScaledCompassPoints(len("XMAS"))
	return runPart(gridPoints, contents, scaledCompassPoints, []string{"XMAS"})
}

func day4_part2(contents []string) int {
	gridPoints := u.PointsInGrid(len(contents[0]), len(contents))
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

func runPart(gridPoints iter.Seq2[int, u.Point], contents []string, lines [][]u.Point, matches []string) int {
	count := 0
	for _, pos := range gridPoints {
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
