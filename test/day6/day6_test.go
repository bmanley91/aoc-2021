package day6

import (
	"testing"

	"manley.dev/aoc/main/day6"
	"manley.dev/aoc/main/util"
)

func TestDay5Part1(t *testing.T) {
	const expected = 5934
	// Given an input list of fish
	fishList, _ := day6.ParseInputIntoFish("data/test-day-6.txt")

	// When the number of fish after 80 days is calculated
	answer := day6.Part1(fishList)

	// Then the correct answer is returned
	util.AssertAnswer(expected, answer, t)
}

func TestDay5Part2(t *testing.T) {
	const expected = 26984457539
	// Given an input list of fish
	fishList, _ := day6.ParseInputIntoFish("data/test-day-6.txt")

	// When the number of fish after 256 days is calculated
	answer := day6.Part2(fishList)

	// Then the correct answer is returned
	util.AssertAnswer(expected, answer, t)
}
