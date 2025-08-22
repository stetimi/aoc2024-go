package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsPositive(t *testing.T) {
	assert.Equal(t, 5, Abs(5))
}

func TestAbsNegative(t *testing.T) {
	assert.Equal(t, 5, Abs(-5))
}

func TestAbsZero(t *testing.T) {
	assert.Equal(t, 0, Abs(0))
}

func TestSignPositive(t *testing.T) {
	assert.Equal(t, 1, Sign(5))
}

func TestSignNegative(t *testing.T) {
	assert.Equal(t, -1, Sign(-5))
}

func TestSignZero(t *testing.T) {
	assert.Equal(t, 0, Sign(0))
}

func TestScalePoint(t *testing.T) {
	p := Point{X: 2, Y: 3}
	scaled := p.Scale(4)
	assert.Equal(t, Point{X: 8, Y: 12}, scaled)
}

func TestScaledLine(t *testing.T) {
	p := Point{X: 1, Y: 2}
	line := p.ScaledLine(3)
	expected := []Point{{0, 0}, {1, 2}, {2, 4}}
	assert.Equal(t, expected, line)
}

func TestPointsInGrid(t *testing.T) {
	indicesList := make([]int, 4)
	pointsList := make([]Point, 4)
	for i, point := range PointsInGrid(2, 2) {
		indicesList[i] = i
		pointsList[i] = point
	}
	assert.Equal(t, []int{0, 1, 2, 3}, indicesList)
	assert.Equal(t, []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, pointsList)
}

func TestPointDifference(t *testing.T) {
	p1 := Point{-2, 3}
	p2 := Point{1, -1}
	expected := Point{-3, 4}
	result := p1.Difference(p2)
	assert.Equal(t, expected, result)
}

func TestPointBack_North(t *testing.T) {
	p := Point{3, 4}
	expected := Point{3, 5}
	result := p.Back(North)
	assert.Equal(t, expected, result)
}

func TestPointBack_West(t *testing.T) {
	p := Point{3, 4}
	expected := Point{4, 4}
	result := p.Back(West)
	assert.Equal(t, expected, result)
}

func TestPointBack_East(t *testing.T) {
	p := Point{3, 4}
	expected := Point{2, 4}
	result := p.Back(East)
	assert.Equal(t, expected, result)
}

func TestPointBack_South(t *testing.T) {
	p := Point{3, 4}
	expected := Point{3, 3}
	result := p.Back(South)
	assert.Equal(t, expected, result)
}
