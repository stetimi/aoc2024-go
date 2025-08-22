package utils

import (
	"strconv"
	"strings"
)

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func MustParseSeparatedInts(s string, separator string) []int {
	parts := strings.Split(s, separator)
	ints := make([]int, len(parts))
	for i, part := range parts {
		ints[i] = MustParseInt(part)
	}
	return ints
}
