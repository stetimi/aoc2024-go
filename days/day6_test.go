package days

import (
	u "aoc2024-go/utils"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	set "github.com/ugurcsen/gods-generic/sets/linkedhashset"
)

//go:embed day6_test.txt
var day6_contents []byte

func TestDay6(t *testing.T) {
	answers := Day6(day6_contents)
	assert.Equal(t, u.IntAnswers(41, 6), answers)
}

func TestMoveGuardFromStartDay1(t *testing.T) {
	grid, guard := buildGrid(day6_contents)
	guardTrail, result := moveGuardFromStart(&grid, guard, dummyPoint)
	assert.Equal(t, 45, guardTrail.Size())
	assert.Equal(t, Exited, result)
}

func TestMoveGuardGetsIntoLoop(t *testing.T) {
	grid, guard := buildGrid(day6_contents)
	guardTrail, result := moveGuardFromStart(&grid, guard, u.Point{X: 3, Y: 6})
	assert.Equal(t, Looped, result)
	assert.Equal(t, 19, guardTrail.Size())
}

func TestMoveGuardNorthWhenThereIsNoObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardNorth(u.Point{X: 5, Y: 5}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	for y := range 5 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 5, Y: 4 - y}, direction: u.North})
	}
	assert.Equal(t, Exited, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 5, Y: 0}, direction: u.North})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardNorthAtBoundary(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardNorth(u.Point{X: 5, Y: 0}, guardTrail, dummyPoint)
	assert.Equal(t, Exited, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 5, Y: 0}, direction: u.North})
	assert.Equal(t, set.New[guardState](), guardTrail)
}

func TestMoveGuardNorthWhenThereIsObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardNorth(u.Point{X: 2, Y: 6}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 2, Y: 5}, direction: u.North})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 2, Y: 4}, direction: u.North})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 2, Y: 4}, direction: u.North})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardNorthWhenThereIsAnExtraObstacleOnAnEmptyColumn(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardNorth(u.Point{X: 5, Y: 5}, guardTrail, u.Point{X: 5, Y: 2})
	expectedGuardStates := set.New[guardState]()
	for y := range 2 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 5, Y: 4 - y}, direction: u.North})
	}
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guardState{pos: u.Point{X: 5, Y: 3}, direction: u.North}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardNorthWhenThereIsAnExtraObstacleOnANonEmptyColumn(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardNorth(u.Point{X: 1, Y: 9}, guardTrail, u.Point{X: 1, Y: 7})
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 1, Y: 8}, direction: u.North})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guardState{pos: u.Point{X: 1, Y: 8}, direction: u.North}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardWestWhenThereIsNoObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardWest(u.Point{X: 3, Y: 0}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	for x := range 3 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 2 - x, Y: 0}, direction: u.West})
	}
	assert.Equal(t, Exited, result)
	assert.Equal(t, guardState{pos: u.Point{X: 0, Y: 0}, direction: u.West}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardWestAtBoundary(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardWest(u.Point{X: 0, Y: 3}, guardTrail, dummyPoint)
	assert.Equal(t, Exited, result)
	assert.Equal(t, guardState{pos: u.Point{X: 0, Y: 3}, direction: u.West}, guard)
	assert.Equal(t, set.New[guardState](), guardTrail)
}

func TestMoveGuardSouthWhenThereIsNoObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardSouth(u.Point{X: 3, Y: 1}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	for y := range 8 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 3, Y: 2 + y}, direction: u.South})
	}
	assert.Equal(t, Exited, result)
	assert.Equal(t, guardState{pos: u.Point{X: 3, Y: 9}, direction: u.South}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardSouthAtBoundary(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardSouth(u.Point{X: 5, Y: 9}, guardTrail, dummyPoint)
	assert.Equal(t, Exited, result)
	assert.Equal(t, guardState{pos: u.Point{X: 5, Y: 9}, direction: u.South}, guard)
	assert.Equal(t, set.New[guardState](), guardTrail)
}

func TestMoveGuardSouthWhenThereIsObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardSouth(u.Point{X: 7, Y: 1}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 7, Y: 2}, direction: u.South})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 7, Y: 3}, direction: u.South})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guardState{pos: u.Point{X: 7, Y: 3}, direction: u.South}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastWhenThereIsNoObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 7, Y: 2}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 8, Y: 2}, direction: u.East})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 9, Y: 2}, direction: u.East})
	assert.Equal(t, Exited, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 9, Y: 2}, direction: u.East})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastAtBoundary(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 9, Y: 2}, guardTrail, dummyPoint)
	assert.Equal(t, Exited, result)
	assert.Equal(t, guardState{pos: u.Point{X: 9, Y: 2}, direction: u.East}, guard)
	assert.Equal(t, set.New[guardState](), guardTrail)
}

func TestMoveGuardEastWhenThereIsObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 0, Y: 3}, guardTrail, dummyPoint)
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastWhenThereAnExtraObstacleOnAnEmptyRow(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 1, Y: 2}, guardTrail, u.Point{X: 5, Y: 2})
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 2, Y: 2}, direction: u.East})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 3, Y: 2}, direction: u.East})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 4, Y: 2}, direction: u.East})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 4, Y: 2}, direction: u.East})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastWhenThereAnExtraObstacleOnANonEmptyRowAndHitImmediately(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 0, Y: 3}, guardTrail, u.Point{X: 1, Y: 3})
	expectedGuardStates := set.New[guardState]()
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 0, Y: 3}, direction: u.East})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastWhenThereIsAnExtraObstacleBeyondTheNextObstacle(t *testing.T) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := set.New[guardState]()
	guard, result := grid.moveGuardEast(u.Point{X: 0, Y: 3}, guardTrail, u.Point{X: 5, Y: 3})
	expectedGuardStates := set.New[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East})
	assert.Equal(t, OnGrid, result)
	assert.Equal(t, guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East}, guard)
	assert.Equal(t, expectedGuardStates, guardTrail)
}
