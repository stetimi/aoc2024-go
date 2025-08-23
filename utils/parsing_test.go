package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParseCommaSeparatedInts(t *testing.T) {
	result := MustParseSeparatedInts("1,-2,3,4,5", ",")
	expected := []int{1, -2, 3, 4, 5}
	assert.Equal(t, expected, result)
}
