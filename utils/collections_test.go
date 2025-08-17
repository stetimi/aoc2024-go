package utils

import (
	"testing"
)

func TestEmptyFrequencyMap(t *testing.T) {
	m := FrequencyMap([]int{})
	if len(m) != 0 {
		t.Errorf("Expected empty map, got %v", m)
	}
}

func TestSeveralValues(t *testing.T) {
	m := FrequencyMap([]int{1, 2, 1, 1, 3, 2})
	if m[1] != 3 || m[2] != 2 || m[3] != 1 {
		t.Errorf("Expected map with frequencies {1: 3, 2: 2, 3: 1}, got %v", m)
	}
	if len(m) != 3 {
		t.Errorf("Expected 3 frequencies, got %d", len(m))
	}
}
