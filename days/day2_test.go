package days

import "testing"

func TestIsSafeIncreasing(t *testing.T) {
	runIsSafeTest(t, "Increasing", []int{1, 2, 4, 6}, true)
}

func TestIsSafeDecreasing(t *testing.T) {
	runIsSafeTest(t, "Decreasing", []int{9, 7, 6, 3, 1}, true)
}

func TestIsSafeNotStrictlyMonotonic(t *testing.T) {
	runIsSafeTest(t, "NotStrictlyMonotonic", []int{1, 3, 3, 4}, false)
}

func TestIsSafeLargeDiff(t *testing.T) {
	runIsSafeTest(t, "LargeDiff", []int{1, 3, 7}, false)
}

func TestIsSafeZigzag(t *testing.T) {
	runIsSafeTest(t, "Zigzag", []int{1, 2, 1, 2}, false)
}

func runIsSafeTest(t *testing.T, name string, input []int, wanted bool) {
	t.Helper()
	result := isSafe(input)
	if result != wanted {
		t.Errorf("%s: isSafe(%v) = %v, wanted %v", name, input, result, wanted)
	}
}
