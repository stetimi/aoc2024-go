package days

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed day5_test.txt
var day5_contents []byte

func TestDay5(t *testing.T) {
	answer := Day5(day5_contents)
	assert.Equal(t, "143", answer.Part1)
	assert.Equal(t, "123", answer.Part2)
}

func TestParseOrderRulesAndUpdates(t *testing.T) {
	input := []byte("47|53\n12|47\n12,47,53\n53,12,47")
	result := parseOrderRulesAndUpdates(input)

	assert.Equal(t, []orderingRule{{before: 47, after: 53}, {before: 12, after: 47}}, result.rules)
	assert.Equal(t, [][]int{{12, 47, 53}, {53, 12, 47}}, result.updates)
}
