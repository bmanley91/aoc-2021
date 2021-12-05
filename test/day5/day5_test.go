package day5

import (
	"testing"

	"manley.dev/aoc/main/day5"
	"manley.dev/aoc/main/util"
)

func TestDay5Part1(t *testing.T) {
	const expected = 5
	// Given sample data
	lines, _ := day5.ParseInputIntoLines("data/test-day-5.txt")

	// When the number of overlapping points is measured
	answer := day5.Part1(lines)

	// Then the correct response is given
	util.AssertAnswer(expected, answer, t)
}
func TestDay5Part2(t *testing.T) {
	const expected = 12
	// Given sample data
	lines, _ := day5.ParseInputIntoLines("data/test-day-5.txt")

	// When the number of overlapping points is measured
	answer := day5.Part2(lines)

	// Then the correct response is given
	util.AssertAnswer(expected, answer, t)
}
