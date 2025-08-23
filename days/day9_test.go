package days

import (
	u "aoc2024-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day9_contents = []byte("2333133121414131402")

func TestDay9(t *testing.T) {
	answers := Day9(day9_contents)
	assert.Equal(t, u.Part1OnlyIntAnswers(1928), answers)
}
