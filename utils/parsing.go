package utils

import (
	"strings"

	"github.com/go-softwarelab/common/pkg/must"
)

func MustParseSeparatedInts(s string, separator string) []int {
	parts := strings.Split(s, separator)
	ints := make([]int, len(parts))
	for i, part := range parts {
		ints[i] = must.ConvertToIntFromString(part)
	}
	return ints
}
