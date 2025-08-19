package utils

import (
	"iter"
)

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

func ScalePoint(p Point, scale int) Point {
	return Point{
		X: p.X * scale,
		Y: p.Y * scale,
	}
}

func ScaledLine(p Point, len int) []Point {
	points := make([]Point, len)
	for i := range len {
		points[i] = ScalePoint(p, i)
	}
	return points
}

func AllCompassPoints() []Point {
	return []Point{
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 1, Y: -1},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
		{X: -1, Y: -1},
	}
}

func PointsInGrid(width, height int) iter.Seq2[int, Point] {
	return func(yield func(int, Point) bool) {
		i := 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !yield(i, Point{X: x, Y: y}) {
					return
				}
				i++
			}
		}
	}
}
