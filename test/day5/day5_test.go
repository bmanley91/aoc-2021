package day5

import (
	"testing"

	"manley.dev/aoc/main/day5"
	"manley.dev/aoc/main/util"
)

func TestDay5Part1(t *testing.T) {
	const expected = 5
	// Given sample data

	// When the number of overlapping points is measured
	answer := day5.Part1()

	// Then the correct response is given
	util.AssertAnswer(expected, answer, t)
}
