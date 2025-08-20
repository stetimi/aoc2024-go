package days

import (
	u "aoc2024-go/utils"
	"bytes"

	mapset "github.com/deckarep/golang-set/v2"
)

type grid struct {
	width              int
	rows               [][]byte
	obstaclesPerRow    map[int][]int
	obstaclesPerColumn map[int][]int
}

type guardState struct {
	pos       u.Point
	direction u.Direction4
}

func Day6(contents []byte) u.Answers {
	grid, guard := buildGrid(contents)
	guardTrail := mapset.NewSet[guardState]()
	guardTrail.Add(guard)
	onGrid := true
	for onGrid {
		guard, onGrid = grid.moveGuard(guard, &guardTrail)
		guard = guardState{pos: guard.pos, direction: guard.direction.TurnRight()}
	}
	part1 := countUniquePositions(guardTrail)
	return u.IntAnswers(part1, 0)
}

func (grid *grid) moveGuard(guard guardState, guardTrail *mapset.Set[guardState]) (guardState, bool) {
	switch guard.direction {
	case u.North:
		return grid.moveGuardNorth(guard, guardTrail)
	case u.East:
		return grid.moveGuardEast(guard, guardTrail)
	case u.South:
		return grid.moveGuardSouth(guard, guardTrail)
	case u.West:
		return grid.moveGuardWest(guard, guardTrail)
	}
	panic("Unknown direction")
}

func countUniquePositions(guardTrail mapset.Set[guardState]) int {
	positions := mapset.NewSet[u.Point]()
	for guardState := range guardTrail.Iter() {
		positions.Add(guardState.pos)
	}
	return positions.Cardinality()
}

func (grid *grid) moveGuardNorth(guard guardState, guardTrail *mapset.Set[guardState]) (guardState, bool) {
	obstacles := grid.obstaclesPerColumn[guard.pos.X]
	obstacleToNorth := findFirstObstacleToNorthOrWest(obstacles, guard.pos.Y)
	guard = guardStatesMovingNorth(guard, guardTrail, obstacleToNorth)
	return guard, obstacleToNorth != -1
}

func (grid *grid) moveGuardEast(guard guardState, guardTrail *mapset.Set[guardState]) (guardState, bool) {
	obstacles := grid.obstaclesPerRow[guard.pos.Y]
	obstacleToEast := findFirstObstacleToSouthOrEast(obstacles, guard.pos.X, grid.width)
	guard = guardStatesMovingEast(guard, guardTrail, obstacleToEast)
	return guard, obstacleToEast != grid.width
}

func (grid *grid) moveGuardSouth(guard guardState, guardTrail *mapset.Set[guardState]) (guardState, bool) {
	obstacles := grid.obstaclesPerColumn[guard.pos.X]
	obstacleToSouth := findFirstObstacleToSouthOrEast(obstacles, guard.pos.Y, len(grid.rows))
	guard = guardStatesMovingSouth(guard, guardTrail, obstacleToSouth)
	return guard, obstacleToSouth != len(grid.rows)
}

func (grid *grid) moveGuardWest(guard guardState, guardTrail *mapset.Set[guardState]) (guardState, bool) {
	obstacles := grid.obstaclesPerRow[guard.pos.Y]
	obstacleToWest := findFirstObstacleToNorthOrWest(obstacles, guard.pos.X)
	guard = guardStatesMovingWest(guard, guardTrail, obstacleToWest)
	return guard, obstacleToWest != -1
}

func findFirstObstacleToNorthOrWest(obstacles []int, p int) int {
	for i := len(obstacles) - 1; i >= 0; i-- {
		if obstacles[i] < p {
			return obstacles[i]
		}
	}
	return -1
}

func findFirstObstacleToSouthOrEast(obstacles []int, p int, endIndex int) int {
	for _, obstacle := range obstacles {
		if obstacle > p {
			return obstacle
		}
	}
	return endIndex
}

func guardStatesMovingNorth(guard guardState, guardTrail *mapset.Set[guardState], endY int) guardState {
	for y := guard.pos.Y - 1; y > endY; y-- {
		guard = guardState{pos: u.Point{X: guard.pos.X, Y: y}, direction: guard.direction}
		(*guardTrail).Add(guard)
	}
	return guard
}

func guardStatesMovingSouth(guard guardState, guardTrail *mapset.Set[guardState], endY int) guardState {
	for y := guard.pos.Y + 1; y < endY; y++ {
		guard = guardState{pos: u.Point{X: guard.pos.X, Y: y}, direction: guard.direction}
		(*guardTrail).Add(guard)
	}
	return guard
}

func guardStatesMovingWest(guard guardState, guardTrail *mapset.Set[guardState], endX int) guardState {
	for x := guard.pos.X - 1; x > endX; x-- {
		guard = guardState{pos: u.Point{X: x, Y: guard.pos.Y}, direction: guard.direction}
		(*guardTrail).Add(guard)
	}
	return guard
}

func guardStatesMovingEast(guard guardState, guardTrail *mapset.Set[guardState], endX int) guardState {
	for x := guard.pos.X + 1; x < endX; x++ {
		guard = guardState{pos: u.Point{X: x, Y: guard.pos.Y}, direction: guard.direction}
		(*guardTrail).Add(guard)
	}
	return guard
}

func buildGrid(contents []byte) (grid, guardState) {
	lines := bytes.Split(contents, []byte("\n"))
	width := len(lines[0])
	height := len(lines)
	rows := make([][]byte, len(lines))
	var guard u.Point
	obstaclesPerRow := make(map[int][]int)
	obstaclesPerColumn := make(map[int][]int)
	for y := range height {
		obstaclesPerRow[y] = []int{}
	}
	for x := range width {
		obstaclesPerColumn[x] = []int{}
	}

	for y, line := range lines {
		rows[y] = line
		for x, cell := range line {
			switch cell {
			case '#':
				obstaclesPerRow[y] = append(obstaclesPerRow[y], x)
				obstaclesPerColumn[x] = append(obstaclesPerColumn[x], y)
			case '^':
				guard = u.Point{X: x, Y: y}
				line[x] = '.'
			}
		}
	}
	initialGrid := grid{
		width:              width,
		rows:               rows,
		obstaclesPerRow:    obstaclesPerRow,
		obstaclesPerColumn: obstaclesPerColumn,
	}
	guardState := guardState{pos: guard, direction: u.North}
	return initialGrid, guardState
}
