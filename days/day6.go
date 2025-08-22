package days

import (
	u "aoc2024-go/utils"
	"bytes"

	set "github.com/ugurcsen/gods-generic/sets/linkedhashset"
)

type grid struct {
	width              int
	height             int
	obstaclesPerRow    map[int][]int
	obstaclesPerColumn map[int][]int
}

type guardState struct {
	pos       u.Point
	direction u.Direction4
}

type guardMovementResult int

const (
	Exited guardMovementResult = iota
	OnGrid
	Looped
)

var dummyPoint = u.Point{X: 1_000_000, Y: 1_000_000}

func Day6(contents []byte) u.Answers {
	grid, guard := buildGrid(contents)
	guardTrail, _ := moveGuardFromStart(&grid, guard, dummyPoint)
	part1 := uniquePositions(guardTrail).Size()
	part2 := part2(&grid, guardTrail, guard)
	return u.IntAnswers(part1, part2)
}

func part2(grid *grid, guardTrail *set.Set[guardState], _ guardState) int {
	testedPositions := set.New[u.Point]()
	loopCount := 0
	for _, trailPoint := range guardTrail.Values() {
		if testedPositions.Contains(trailPoint.pos) {
			continue
		}
		testedPositions.Add(trailPoint.pos)
		direction := trailPoint.direction
		guardStart := guardState{pos: trailPoint.pos.Back(direction), direction: direction}
		_, result := moveGuardFromStart(grid, guardStart, trailPoint.pos)
		if result == Looped {
			loopCount += 1
		}
	}
	return loopCount
}

func moveGuardFromStart(grid *grid, guard guardState, extraObstacle u.Point) (*set.Set[guardState], guardMovementResult) {
	guardTrail := set.New[guardState]()
	guardTrail.Add(guard)
	result := grid.moveGuardRepeatedly(guard, guardTrail, extraObstacle)
	return guardTrail, result
}

func (grid *grid) moveGuardRepeatedly(guard guardState, guardTrail *set.Set[guardState], extraObstacle u.Point) guardMovementResult {
	result := OnGrid
	for result == OnGrid {
		guard, result = grid.moveGuard(guard, guardTrail, extraObstacle)
		guard = guardState{pos: guard.pos, direction: guard.direction.TurnRight()}
	}
	return result
}

func (grid *grid) moveGuard(guard guardState, guardTrail *set.Set[guardState], extraObstacle u.Point) (guardState, guardMovementResult) {
	switch guard.direction {
	case u.North:
		return grid.moveGuardNorth(guard.pos, guardTrail, extraObstacle)
	case u.East:
		return grid.moveGuardEast(guard.pos, guardTrail, extraObstacle)
	case u.South:
		return grid.moveGuardSouth(guard.pos, guardTrail, extraObstacle)
	case u.West:
		return grid.moveGuardWest(guard.pos, guardTrail, extraObstacle)
	}
	panic("Unknown direction")
}

func uniquePositions(guardTrail *set.Set[guardState]) *set.Set[u.Point] {
	positions := set.New[u.Point]()
	guardTrail.Each(func(_ int, guardState guardState) {
		positions.Add(guardState.pos)
	})
	return positions
}

func (grid *grid) moveGuardNorth(pos u.Point, guardTrail *set.Set[guardState], extraObstacle u.Point) (guardState, guardMovementResult) {
	obstacles := grid.obstaclesPerColumn[pos.X]
	extraObstacleY := dummyPoint.Y
	if extraObstacle.X == pos.X {
		extraObstacleY = extraObstacle.Y
	}
	obstacleToNorth := findFirstObstacleToNorthOrWest(obstacles, extraObstacleY, pos.Y)
	return guardStatesMovingNorth(pos, guardTrail, obstacleToNorth)
}

func (grid *grid) moveGuardEast(pos u.Point, guardTrail *set.Set[guardState], extraObstacle u.Point) (guardState, guardMovementResult) {
	obstacles := grid.obstaclesPerRow[pos.Y]
	extraObstacleX := 1_000_000
	if extraObstacle.Y == pos.Y {
		extraObstacleX = extraObstacle.X
	}
	obstacleToEast := findFirstObstacleToSouthOrEast(obstacles, extraObstacleX, pos.X, grid.width)
	return guardStatesMovingEast(pos, guardTrail, obstacleToEast, grid.width)
}

func (grid *grid) moveGuardSouth(pos u.Point, guardTrail *set.Set[guardState], extraObstacle u.Point) (guardState, guardMovementResult) {
	obstacles := grid.obstaclesPerColumn[pos.X]
	extraObstacleY := 1_000_000
	if extraObstacle.X == pos.X {
		extraObstacleY = extraObstacle.Y
	}
	obstacleToSouth := findFirstObstacleToSouthOrEast(obstacles, extraObstacleY, pos.Y, grid.height)
	return guardStatesMovingSouth(pos, guardTrail, obstacleToSouth, grid.height)
}

func (grid *grid) moveGuardWest(pos u.Point, guardTrail *set.Set[guardState], extraObstacle u.Point) (guardState, guardMovementResult) {
	obstacles := grid.obstaclesPerRow[pos.Y]
	extraObstacleX := 1_000_000
	if extraObstacle.Y == pos.Y {
		extraObstacleX = extraObstacle.X
	}
	obstacleToWest := findFirstObstacleToNorthOrWest(obstacles, extraObstacleX, pos.X)
	return guardStatesMovingWest(pos, guardTrail, obstacleToWest)
}

func findFirstObstacleToNorthOrWest(obstacles []int, extraObstacleCoord int, p int) int {
	firstObstacle := -1
	for i := len(obstacles) - 1; i >= 0; i-- {
		if obstacles[i] < p {
			firstObstacle = obstacles[i]
			break
		}
	}
	if extraObstacleCoord < p && extraObstacleCoord > firstObstacle {
		return extraObstacleCoord
	}
	return firstObstacle
}

func findFirstObstacleToSouthOrEast(obstacles []int, extraObstacleCoord int, p int, endIndex int) int {
	firstObstacle := endIndex
	for _, obstacle := range obstacles {
		if obstacle > p {
			firstObstacle = obstacle
			break
		}
	}
	if extraObstacleCoord > p && extraObstacleCoord < firstObstacle {
		return extraObstacleCoord
	}
	return firstObstacle
}

func guardStatesMovingNorth(pos u.Point, guardTrail *set.Set[guardState], endY int) (guardState, guardMovementResult) {
	guard := guardState{pos: pos, direction: u.North}
	for y := guard.pos.Y - 1; y > endY; y-- {
		guard = guardState{pos: u.Point{X: guard.pos.X, Y: y}, direction: u.North}
		if guardTrail.Contains(guard) {
			return guard, Looped
		}
		guardTrail.Add(guard)
	}
	result := OnGrid
	if guard.pos.Y == 0 {
		result = Exited
	}
	return guard, result
}

func guardStatesMovingSouth(pos u.Point, guardTrail *set.Set[guardState], endY int, southernEndIndex int) (guardState, guardMovementResult) {
	guard := guardState{pos: pos, direction: u.South}
	for y := guard.pos.Y + 1; y < endY; y++ {
		guard = guardState{pos: u.Point{X: pos.X, Y: y}, direction: u.South}
		if guardTrail.Contains(guard) {
			return guard, Looped
		}
		guardTrail.Add(guard)
	}
	result := OnGrid
	if guard.pos.Y == southernEndIndex-1 {
		result = Exited
	}
	return guard, result
}

func guardStatesMovingWest(pos u.Point, guardTrail *set.Set[guardState], endX int) (guardState, guardMovementResult) {
	guard := guardState{pos: pos, direction: u.West}
	for x := pos.X - 1; x > endX; x-- {
		guard = guardState{pos: u.Point{X: x, Y: pos.Y}, direction: u.West}
		if guardTrail.Contains(guard) {
			return guard, Looped
		}
		guardTrail.Add(guard)
	}
	result := OnGrid
	if guard.pos.X == 0 {
		result = Exited
	}
	return guard, result
}

func guardStatesMovingEast(pos u.Point, guardTrail *set.Set[guardState], endX int, easternEndIndex int) (guardState, guardMovementResult) {
	guard := guardState{pos: pos, direction: u.East}
	for x := pos.X + 1; x < endX; x++ {
		guard = guardState{pos: u.Point{X: x, Y: pos.Y}, direction: guard.direction}
		if guardTrail.Contains(guard) {
			return guard, Looped
		}
		guardTrail.Add(guard)
	}
	result := OnGrid
	if guard.pos.X == easternEndIndex-1 {
		result = Exited
	}
	return guard, result
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
			}
		}
	}
	initialGrid := grid{
		width:              width,
		height:             height,
		obstaclesPerRow:    obstaclesPerRow,
		obstaclesPerColumn: obstaclesPerColumn,
	}
	guardState := guardState{pos: guard, direction: u.North}
	return initialGrid, guardState
}
