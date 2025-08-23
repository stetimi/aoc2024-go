package days

import (
	u "aoc2024-go/utils"
	_ "embed"
	"testing"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/stretchr/testify/assert"
)

//go:embed day10_test.txt
var day10_contents []byte

func TestDay10(t *testing.T) {
	answers := Day10(day10_contents)
	assert.Equal(t, u.Part1OnlyIntAnswers(36), answers)
}

func TestSmallSample(t *testing.T) {
	contents := []byte(`...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`)
	grid, width := readGrid(contents)
	result := day10Part1(grid, width)
	assert.Equal(t, 2, result)
}

func TestAnotherSmallSample(t *testing.T) {
	contents := []byte(`..90..9
...1.98
...2..7
6543456
765.987
876....
987....`)
	grid, width := readGrid(contents)
	result := day10Part1(grid, width)
	assert.Equal(t, 4, result)
}

func TestLargerSample(t *testing.T) {
	contents := []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)
	grid, width := readGrid(contents)
	result := day10Part1(grid, width)
	assert.Equal(t, 36, result)
}

func TestDepthFirstSearch(t *testing.T) {
	graph := map[int][]int{
		1: {2, 5},
		2: {3},
		5: {4},
	}
	next := func(n int) []int {
		return graph[n]
	}
	isTarget := func(n int) bool {
		return n >= 3
	}
	result := seq.Collect(depthFirstSearch(1, next, isTarget))
	expected := []int{3, 5, 4}
	assert.ElementsMatch(t, expected, result, "DFS traversal should find all target nodes in the graph")
}

func TestNextSteps(t *testing.T) {
	t.Run("All 4 points available", func(t *testing.T) {
		// 827
		// 219
		// 525
		grid := []byte("827219525")
		width := 3
		result := nextSteps(grid, width, 4)
		assert.ElementsMatch(t, []int{1, 3, 7}, result)
	})
	t.Run("No point to north", func(t *testing.T) {
		// 232
		// 523
		grid := []byte("232523")
		width := 3
		result := nextSteps(grid, width, 4)
		assert.Equal(t, []int{1, 5}, result)
	})
	t.Run("No point to east", func(t *testing.T) {
		// 234
		// 423
		grid := []byte("234423")
		width := 3
		result := nextSteps(grid, width, 5)
		assert.Equal(t, []int{2}, result)
	})
	t.Run("No point to west", func(t *testing.T) {
		// 234
		// 123
		grid := []byte("234123")
		width := 3
		result := nextSteps(grid, width, 3)
		assert.Equal(t, []int{0, 4}, result)
	})
	t.Run("No point to south", func(t *testing.T) {
		// 234
		// 123
		grid := []byte("234123")
		width := 3
		result := nextSteps(grid, width, 4)
		assert.Equal(t, []int{1, 5}, result)
	})
}
