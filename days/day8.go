package days

import (
	u "aoc2024-go/utils"
	"bytes"

	"iter"

	"github.com/go-softwarelab/common/pkg/seq"
	set "github.com/ugurcsen/gods-generic/sets/hashset"
)

func Day8(contents []byte) u.Answers {
	rows := bytes.Split(contents, []byte("\n"))
	width := len(rows[0])
	height := len(rows)
	antennaPositions := parseAntennaPositions(rows)
	part1 := day8Parts(antennaPositions, width, height, seq.Of(1))
	part2 := day8Parts(antennaPositions, width, height, seq.Range(0, 1_000_000))
	return u.IntAnswers(part1, part2)
}

func day8Parts(antennaPositions map[byte][]u.Point, width, height int, scales iter.Seq[int]) int {
	antiNodePoints := set.New[u.Point]()
	for _, antennaPositions := range antennaPositions {
		for antiNodePoint := range antiNodes(pairsOf(antennaPositions), width, height, scales) {
			antiNodePoints.Add(antiNodePoint)
		}
	}
	return antiNodePoints.Size()
}

func parseAntennaPositions(rows [][]byte) map[byte][]u.Point {
	antennaPositions := make(map[byte][]u.Point)
	for y, row := range rows {
		for x, b := range row {
			if b != '.' {
				antennaPositions[b] = append(antennaPositions[b], u.Point{X: x, Y: y})
			}
		}
	}
	return antennaPositions
}

func antiNodes(pairs iter.Seq[[2]u.Point], width, height int, scales iter.Seq[int]) iter.Seq[u.Point] {
	return func(yield func(u.Point) bool) {
		for pair := range pairs {
			dist := pair[1].Difference(pair[0])
			for scale := range scales {
				antiNode := pair[0].AddScaled(dist, -scale)
				if !antiNode.IsInside(width, height) {
					break
				}
				if !yield(antiNode) {
					return
				}
			}
			for scale := range scales {
				antiNode := pair[1].AddScaled(dist, scale)
				if !antiNode.IsInside(width, height) {
					break
				}
				if !yield(antiNode) {
					return
				}
			}
		}
	}
}

func pairsOf(points []u.Point) iter.Seq[[2]u.Point] {
	return func(yield func([2]u.Point) bool) {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				if !yield([2]u.Point{points[i], points[j]}) {
					return
				}
			}
		}
	}
}
