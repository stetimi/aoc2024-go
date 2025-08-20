package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParseIntValid(t *testing.T) {
	assert.Equal(t, 42, MustParseInt("42"))
}

func TestMustParseIntZero(t *testing.T) {
	assert.Equal(t, 0, MustParseInt("0"))
}

func TestMustParseIntNegative(t *testing.T) {
	assert.Equal(t, -123, MustParseInt("-123"))
}

func TestMustParseCommaSeparatedInts(t *testing.T) {
	result := MustParseCommaSeparatedInts("1,-2,3,4,5")
	expected := []int{1, -2, 3, 4, 5}
	assert.Equal(t, expected, result)
}
