package days

import (
	u "aoc2024-go/utils"
	_ "embed"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

//go:embed day6_test.txt
var day6_contents []byte

func TestDay6(t *testing.T) {
	answer := Day6(day6_contents)
	assert.Equal(t, "41", answer.Part1)
	// assert.Equal(t, "1686", answer.Part2)
}

func TestMoveGuardNorthWhenThereIsNoObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(5, 5, u.North)
	guard, onGrid := grid.moveGuardNorth(guard, &guardTrail)
	assert.False(t, onGrid)
	expectedGuardStates := mapset.NewSet[guardState]()
	for y := range 5 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 5, Y: y}, direction: u.North})
	}
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 5, Y: 0}, direction: u.North})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardNorthAtBoundary(t *testing.T) {
	grid, guardTrail, guard := fixture(5, 0, u.North)
	guard, onGrid := grid.moveGuardNorth(guard, &guardTrail)
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 5, Y: 0}, direction: u.North})
	assert.Equal(t, mapset.NewSet[guardState](), guardTrail)
}

func TestMoveGuardNorthWhenThereIsObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(2, 6, u.North)
	guard, onGrid := grid.moveGuardNorth(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 2, Y: 5}, direction: u.North})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 2, Y: 4}, direction: u.North})
	assert.True(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 2, Y: 4}, direction: u.North})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardWestWhenThereIsNoObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(3, 0, u.West)
	guard, onGrid := grid.moveGuardWest(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	for x := range 3 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: x, Y: 0}, direction: u.West})
	}
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 0, Y: 0}, direction: u.West})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardWestAtBoundary(t *testing.T) {
	grid, guardTrail, guard := fixture(0, 3, u.West)
	guard, onGrid := grid.moveGuardWest(guard, &guardTrail)
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 0, Y: 3}, direction: u.West})
	assert.Equal(t, mapset.NewSet[guardState](), guardTrail)
}

func TestMoveGuardSouthWhenThereIsNoObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(3, 1, u.South)
	guard, onGrid := grid.moveGuardSouth(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	for y := range 8 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 3, Y: 2 + y}, direction: u.South})
	}
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 3, Y: 9}, direction: u.South})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardSouthAtBoundary(t *testing.T) {
	grid, guardTrail, guard := fixture(5, 9, u.South)
	guard, onGrid := grid.moveGuardSouth(guard, &guardTrail)
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 5, Y: 9}, direction: u.South})
	assert.Equal(t, mapset.NewSet[guardState](), guardTrail)
}

func TestMoveGuardSouthWhenThereIsObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(7, 1, u.South)
	guard, onGrid := grid.moveGuardSouth(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 7, Y: 2}, direction: u.South})
	expectedGuardStates.Add(guardState{pos: u.Point{X: 7, Y: 3}, direction: u.South})
	assert.True(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 7, Y: 3}, direction: u.South})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastWhenThereIsNoObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(7, 2, u.East)
	guard, onGrid := grid.moveGuardEast(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	for x := range 2 {
		expectedGuardStates.Add(guardState{pos: u.Point{X: 8 + x, Y: 2}, direction: u.East})
	}
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 9, Y: 2}, direction: u.East})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func TestMoveGuardEastAtBoundary(t *testing.T) {
	grid, guardTrail, guard := fixture(9, 2, u.East)
	guard, onGrid := grid.moveGuardEast(guard, &guardTrail)
	assert.False(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 9, Y: 2}, direction: u.East})
	assert.Equal(t, mapset.NewSet[guardState](), guardTrail)
}

func TestMoveGuardEastWhenThereIsObstacle(t *testing.T) {
	grid, guardTrail, guard := fixture(0, 3, u.East)
	guard, onGrid := grid.moveGuardEast(guard, &guardTrail)
	expectedGuardStates := mapset.NewSet[guardState]()
	expectedGuardStates.Add(guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East})
	assert.True(t, onGrid)
	assert.Equal(t, guard, guardState{pos: u.Point{X: 1, Y: 3}, direction: u.East})
	assert.Equal(t, expectedGuardStates, guardTrail)
}

func fixture(x int, y int, guardDirection u.Direction4) (grid, mapset.Set[guardState], guardState) {
	grid, _ := buildGrid(day6_contents)
	guardTrail := mapset.NewSet[guardState]()
	return grid, guardTrail, guardState{pos: u.Point{X: x, Y: y}, direction: guardDirection}
}
