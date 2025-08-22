package days

import (
	u "aoc2024-go/utils"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	set "github.com/ugurcsen/gods-generic/sets/hashset"
)

//go:embed day7_test.txt
var day7_contents []byte

func TestDay7(t *testing.T) {
	answers := Day7(day7_contents)
	assert.Equal(t, u.IntAnswers(3749, 11387), answers)
}

func TestReadEquation(t *testing.T) {
	assert.Equal(t, equation{target: 994, numbers: []int{10, 2, 3}}, readEquation("994: 10 2 3"))
}

func TestResultsForNextNumberNotHittingTarget(t *testing.T) {
	results, evalResult := resultsForNextNumber(50, set.New(5, 8), []intOp{addOp, mulOp}, 2, true)
	assert.Equal(t, set.New(7, 10, 16), results)
	assert.Equal(t, MoreNumbersNeeded, evalResult)
}

func TestResultsForNextNumberHittingTargetButThereAreStillNumbersToGo(t *testing.T) {
	results, evalResult := resultsForNextNumber(15, set.New(5, 8), []intOp{addOp, mulOp}, 3, false)
	assert.Equal(t, set.New(8, 11, 15), results)
	assert.Equal(t, MoreNumbersNeeded, evalResult)
}

func TestResultsForNextNumberHittingTargetAndThereAreNoMoreNumbersToGo(t *testing.T) {
	results, evalResult := resultsForNextNumber(15, set.New(5, 8), []intOp{addOp, mulOp}, 3, true)
	assert.Nil(t, results)
	assert.Equal(t, Calibrated, evalResult)
}

func TestConcatOp(t *testing.T) {
	assert.Equal(t, 102, concatOp(10, 2))
	assert.Equal(t, 98161, concatOp(98, 161))
}
