package utils

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
