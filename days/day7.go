package days

import (
	u "aoc2024-go/utils"
	"fmt"
	"strings"

	"github.com/go-softwarelab/common/pkg/must"
	set "github.com/ugurcsen/gods-generic/sets/hashset"
)

type evaluationResult int

const (
	Impossible evaluationResult = iota
	MoreNumbersNeeded
	Calibrated
)

type equation struct {
	target  int
	numbers []int
}

type intOp = func(int, int) int

var addOp intOp = func(a, b int) int { return a + b }
var mulOp intOp = func(a, b int) int { return a * b }
var concatOp intOp = func(a, b int) int {
	concat_str := fmt.Sprintf("%d%d", a, b)
	return must.ConvertToIntFromString(concat_str)
}

func Day7(contents []byte) u.Answers {
	equations := readEquations(contents)
	part1 := runDay7Part(equations, []intOp{addOp, mulOp})
	part2 := runDay7Part(equations, []intOp{addOp, mulOp, concatOp})
	return u.IntAnswers(part1, part2)
}

func runDay7Part(equations []equation, ops []intOp) int {
	result := 0
	for _, eq := range equations {
		if eq.canBeCalibrated(ops) {
			result += eq.target
		}
	}
	return result
}

func (eq *equation) canBeCalibrated(ops []intOp) bool {
	results := set.New[int]()
	results.Add(eq.numbers[0])
	for i, number := range eq.numbers[1:] {
		isLast := i == len(eq.numbers)-2
		var evalResult evaluationResult
		results, evalResult = resultsForNextNumber(eq.target, results, ops, number, isLast)
		switch evalResult {
		case Calibrated:
			return true
		case Impossible:
			return false
		}
	}
	return false
}

func resultsForNextNumber(
	target int,
	results *set.Set[int],
	ops []intOp,
	number int,
	noMoreNumbers bool) (*set.Set[int], evaluationResult) {
	newResults := set.New[int]()
	for _, result := range results.Values() {
		for _, op := range ops {
			newResult := op(result, number)
			if newResult <= target {
				newResults.Add(newResult)
				if noMoreNumbers && newResult == target {
					return nil, Calibrated
				}
			}
		}
	}
	return newResults, MoreNumbersNeeded
}

func readEquations(contents []byte) []equation {
	var equations []equation
	for _, line := range strings.Split(string(contents), "\n") {
		eq := readEquation(line)
		equations = append(equations, eq)
	}
	return equations
}

func readEquation(line string) equation {
	parts := strings.Split(line, ": ")
	target := must.ConvertToIntFromString(parts[0])
	numbers := u.MustParseSeparatedInts(parts[1], " ")
	return equation{target: target, numbers: numbers}
}
