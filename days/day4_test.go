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

func TestNoMatchAtEdge(t *testing.T) {
	pos := u.Point{X: 0, Y: 0}
	assert.Equal(t, 0, countMatches(lines, pos, "XMAS"))
}

func TestMatchAtEdge(t *testing.T) {
	pos := u.Point{X: 0, Y: 4}
	assert.False(t, hasMatchAlongDirection(lines, pos, u.Point{X: -1, Y: 0}, "XMAS"))
	assert.True(t, hasMatchAlongDirection(lines, pos, u.Point{X: 1, Y: 0}, "XMAS"))
	assert.Equal(t, 1, countMatches(lines, pos, "XMAS"))
}
