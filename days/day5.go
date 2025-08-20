package days

import (
	u "aoc2024-go/utils"
	"slices"
	"strings"

	"github.com/dominikbraun/graph"
)

type orderingRule struct {
	before int
	after  int
}

type day5Input struct {
	rules   []orderingRule
	updates [][]int
}

func Day5(contents []byte) u.Answers {
	input := parseOrderRulesAndUpdates(contents)
	part1, part2 := 0, 0
	for _, update := range input.updates {
		sorted := sortUpdate(update, input.rules)
		if slices.Equal(sorted, update) {
			part1 += middle(update)
		} else {
			part2 += middle(sorted)
		}
	}
	return u.IntAnswers(part1, part2)
}

func sortUpdate(pages []int, rules []orderingRule) []int {
	g := graph.New(graph.IntHash, graph.Directed())
	for _, rule := range rules {
		if slices.Contains(pages, rule.before) && slices.Contains(pages, rule.after) {
			_ = g.AddVertex(rule.before)
			_ = g.AddVertex(rule.after)
			err := g.AddEdge(rule.before, rule.after)
			if err != nil {
				panic(err)
			}
		}
	}
	sorted, _ := graph.TopologicalSort(g)
	return sorted
}

func middle(pages []int) int {
	middleIndex := len(pages) / 2
	return pages[middleIndex]
}

func parseOrderRulesAndUpdates(contents []byte) day5Input {
	lines := strings.Split(string(contents), "\n")
	rules := make([]orderingRule, 0)
	updates := make([][]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			splitByPipe := strings.Split(line, "|")
			rule := orderingRule{before: u.MustParseInt(splitByPipe[0]), after: u.MustParseInt(splitByPipe[1])}
			rules = append(rules, rule)
		} else {
			update := u.MustParseCommaSeparatedInts(line)
			updates = append(updates, update)
		}
	}
	return day5Input{rules: rules, updates: updates}
}
