package utils

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Sign(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return 1
	}
	return 0
}

type Point struct {
	X int
	Y int
}

func AllCompassPoints() []Point {
	return []Point{
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: 1},
		{X: 0, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
		{X: -1, Y: -1},
	}
}

func PointsInGrid(width, height int) []Point {
	points := make([]Point, 0, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			points = append(points, Point{X: x, Y: y})
		}
	}
	return points
}
