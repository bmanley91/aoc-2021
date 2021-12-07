package day5

import (
	"testing"

	"manley.dev/aoc/util"
)

func TestDay5Part1(t *testing.T) {
	const expected = 5
	// Given sample data
	lines, _ := ParseInputIntoLines("data/test-day-5.txt")

	// When the number of overlapping points is measured
	answer := Part1(lines)

	// Then the correct response is given
	util.AssertAnswer(expected, answer, t)
}
func TestDay5Part2(t *testing.T) {
	const expected = 12
	// Given sample data
	lines, _ := ParseInputIntoLines("data/test-day-5.txt")

	// When the number of overlapping points is measured
	answer := Part2(lines)

	// Then the correct response is given
	util.AssertAnswer(expected, answer, t)
}
