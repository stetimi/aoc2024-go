package days

import (
	u "aoc2024-go/utils"
	"bytes"
	_ "embed"
	"testing"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/stretchr/testify/assert"
)

//go:embed day8_test.txt
var day8_contents []byte

var smallTestPointsA = []u.Point{
	{X: 4, Y: 3},
	{X: 8, Y: 4},
	{X: 5, Y: 5},
}

func TestDay8(t *testing.T) {
	answers := Day8(day8_contents)
	assert.Equal(t, u.IntAnswers(14, 34), answers)
}

func TestParseAntennaPositions(t *testing.T) {
	points0 := []u.Point{
		{X: 8, Y: 1},
		{X: 5, Y: 2},
		{X: 7, Y: 3},
		{X: 4, Y: 4},
	}
	pointsA := []u.Point{
		{X: 6, Y: 5},
		{X: 8, Y: 8},
		{X: 9, Y: 9},
	}
	rows := bytes.Split([]byte(day8_contents), []byte("\n"))
	result := parseAntennaPositions(rows)
	assert.ElementsMatch(t, points0, result['0'], "Positions for '0' do not match")
	assert.ElementsMatch(t, pointsA, result['A'], "Positions for 'A' do not match")
}

func TestAntiNodesWith1AntinodePerPair(t *testing.T) {
	antiPoints := seq.Collect(antiNodes(pairsOf(smallTestPointsA), 10, 10, seq.Of(1)))
	expectedAntiPoints := []u.Point{
		{X: 3, Y: 1},
		{X: 0, Y: 2},
		{X: 2, Y: 6},
		{X: 6, Y: 7},
	}
	assert.ElementsMatch(t, expectedAntiPoints, antiPoints, "Expected anti-nodes from pointsA")
}

func TestAntiNodesWith1000AntinodePerPair(t *testing.T) {
	nodePoints := []u.Point{
		{X: 0, Y: 0},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
	}
	antiPoints := seq.Collect(seq.Uniq(antiNodes(pairsOf(nodePoints), 10, 10, seq.Range(0, 1_000))))
	expectedAntiPoints := []u.Point{
		{X: 0, Y: 0},
		{X: 5, Y: 0},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
		{X: 6, Y: 2},
		{X: 9, Y: 3},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
		{X: 4, Y: 8},
	}
	assert.ElementsMatch(t, expectedAntiPoints, antiPoints, "Expected anti-nodes from pointsA")
}

func TestPairsOf(t *testing.T) {
	t.Run("single point yields no pairs", func(t *testing.T) {
		singlePoint := []u.Point{{X: 1, Y: 1}}
		pairs := seq.Collect(pairsOf(singlePoint))
		assert.Empty(t, pairs, "Expected no pairs from single point")
	})

	t.Run("two points yield one pair", func(t *testing.T) {
		twoPoints := []u.Point{{X: 1, Y: 1}, {X: 2, Y: 2}}
		pairs := seq.Collect(pairsOf(twoPoints))
		expected := [][2]u.Point{{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
		}}
		assert.Equal(t, expected, pairs, "Expected one pair from two points")
	})

	t.Run("three points yield three pairs", func(t *testing.T) {
		pairs := seq.Collect(pairsOf(smallTestPointsA))
		expected := [][2]u.Point{{
			{X: 4, Y: 3}, {X: 8, Y: 4},
		}, {
			{X: 4, Y: 3}, {X: 5, Y: 5},
		}, {
			{X: 8, Y: 4}, {X: 5, Y: 5},
		}}
		assert.ElementsMatch(t, expected, pairs, "Expected three pairs from three points")
	})
}
