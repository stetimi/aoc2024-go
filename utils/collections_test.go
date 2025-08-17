package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFrequencyMap(t *testing.T) {
	m := FrequencyMap([]int{})
	assert.Empty(t, m)
}

func TestSeveralValues(t *testing.T) {
	m := FrequencyMap([]int{1, 2, 1, 1, 3, 2})
	assert.Equal(t, map[int]int{1: 3, 2: 2, 3: 1}, m)
}
