package days

import (
	u "aoc2024-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lines = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func CountMatchesAtEdgeWhenNone(t *testing.T) {
	pos := u.Point{X: 0, Y: 0}
	assert.Equal(t, 0, countMatches(lines, pos, u.AllCompassPoints(), []string{"XMAS"}))
}

func CountMatchesAtEdge(t *testing.T) {
	pos := u.Point{X: 0, Y: 4}
	assert.Equal(t, 1, countMatches(lines, pos, u.AllCompassPoints(), []string{"XMAS"}))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 18, part1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 9, part2(lines))
}
