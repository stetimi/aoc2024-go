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
