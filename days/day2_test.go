package days

import (
	"fmt"
	"testing"
)

func TestIsSafeIncreasing(t *testing.T) {
	runIsSafeTest(t, "Increasing", []int{1, 2, 4, 6}, -1, true)
}

func TestIsSafeDecreasing(t *testing.T) {
	runIsSafeTest(t, "Decreasing", []int{9, 7, 6, 3, 1}, -1, true)
}

func TestIsSafeNotStrictlyMonotonic(t *testing.T) {
	runIsSafeTest(t, "NotStrictlyMonotonic", []int{1, 3, 3, 4}, -1, false)
}

func TestIsSafeLargeDiff(t *testing.T) {
	runIsSafeTest(t, "LargeDiff", []int{1, 3, 7}, -1, false)
}

func TestIsSafeZigzag(t *testing.T) {
	runIsSafeTest(t, "Zigzag", []int{1, 2, 1, 2}, -1, false)
}

func TestIsSafeIgnoringSingleBadValue(t *testing.T) {
	runIsSafeTest(t, "IsSafeIgnoringSingleBadValue", []int{1, 2, 1, 3}, 2, true)
}

func TestIsNotSafeIgnoringAnyValue(t *testing.T) {
	for i := 0; i <= 4; i++ {
		runIsSafeTest(t, fmt.Sprintf("IsNotSafeIgnoringAnyValue: %d", i), []int{1, 2, 7, 8, 9}, i, false)
	}
}

func runIsSafeTest(t *testing.T, name string, input []int, ignore_index int, wanted bool) {
	t.Helper()
	result := isSafe(input, ignore_index)
	if result != wanted {
		t.Errorf("%s: isSafe(%v) = %v, wanted %v", name, input, result, wanted)
	}
}
