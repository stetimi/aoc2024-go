package days

import (
	u "aoc2024-go/utils"
	"bytes"
	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
	set "github.com/ugurcsen/gods-generic/sets/linkedhashset"
)

func Day10(contents []byte) u.Answers {
	grid, width := readGrid(contents)
	part1 := day10Part1(grid, width)
	part2 := day10Part2(grid, width)
	return u.IntAnswers(part1, part2)
}

func day10Part1(grid []byte, width int) int {
	return sumTrailheadRatings(grid, func() func(pos int) []int { return nextStepsDropVisited(grid, width) })
}

func day10Part2(grid []byte, width int) int {
	next := func(pos int) []int {
		return nextSteps(grid, width, pos)
	}
	return sumTrailheadRatings(grid, func() func(pos int) []int { return next })
}

func sumTrailheadRatings(grid []byte, mkNext func() func(pos int) []int) int {
	isTarget := func(pos int) bool {
		return grid[pos] == '9'
	}
	totalRating := 0
	for trailhead := range findTrailheads(grid) {
		next := mkNext()
		paths := depthFirstSearch(trailhead, next, isTarget)
		score := seq.Count(paths)
		totalRating += score
	}
	return totalRating
}

func nextStepsDropVisited(grid []byte, width int) func(pos int) []int {
	visited := set.New[int]()
	return func(pos int) []int {
		filtered := []int{}
		result := nextSteps(grid, width, pos)
		for _, next := range result {
			if visited.Contains(next) {
				continue
			}
			filtered = append(filtered, next)
			visited.Add(next)
		}
		return filtered
	}
}

func readGrid(contents []byte) (grid []byte, width int) {
	rows := bytes.Split(contents, []byte("\n"))
	grid = bytes.Join(rows, []byte(""))
	width = len(rows[0])
	return
}

func findTrailheads(grid []byte) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range len(grid) {
			if grid[i] == '0' {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func nextSteps(grid []byte, width int, pos int) []int {
	result := []int{}
	height := grid[pos]
	if height == '9' || height == '.' {
		return result
	}
	next_height := height + 1
	column := pos % width
	north := pos - width
	if north >= 0 && grid[north] == next_height {
		result = append(result, north)
	}
	west := pos - 1
	if column > 0 && grid[west] == next_height {
		result = append(result, west)
	}
	east := pos + 1
	if column < width-1 && grid[east] == next_height {
		result = append(result, east)
	}
	south := pos + width
	if south < len(grid) && grid[south] == next_height {
		result = append(result, south)
	}
	return result
}

func depthFirstSearch[T comparable](init T, next func(T) []T, isTarget func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		stack := []T{init}
		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if isTarget(curr) {
				if !yield(curr) {
					return
				}
			}
			nextNodes := next(curr)
			for i := len(nextNodes) - 1; i >= 0; i-- {
				nextNode := nextNodes[i]
				stack = append(stack, nextNode)
			}
		}
	}
}
