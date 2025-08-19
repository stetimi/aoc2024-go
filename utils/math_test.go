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
	scaled := ScalePoint(p, 4)
	assert.Equal(t, Point{X: 8, Y: 12}, scaled)
}

func TestScaledLine(t *testing.T) {
	p := Point{X: 1, Y: 2}
	line := ScaledLine(p, 3)
	expected := []Point{{0, 0}, {1, 2}, {2, 4}}
	assert.Equal(t, expected, line)
}

func TestPointsInGrid(t *testing.T) {
	points := PointsInGrid(2, 2)
	expected := []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
	assert.Equal(t, expected, points)
}
