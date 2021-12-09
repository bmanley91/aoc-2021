package day9

import (
	"testing"

	"manley.dev/aoc/util"
)

func TestDay9Part1(t *testing.T) {
	expected := 15
	// Given an input heightmap
	heights, _ := ParseInput("data/test-day-9.txt")

	// When the total risk score is calculated
	answer := Part1(heights)

	// The the correct response is returned
	util.AssertAnswer(expected, answer, t)
}

// func TestDay9Part2(t *testing.T) {
// 	expected := 1134
// 	// Given an input heightmap
// 	positions, _ := ParseInput("data/test-day-9.txt")

// 	// When the total risk score is calculated
// 	answer := Part2(positions)

// 	// The the correct response is returned
// 	util.AssertAnswer(expected, answer, t)
// }
