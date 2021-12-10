package day10

import (
	"testing"

	"manley.dev/aoc/util"
)

func TestDay10Part1(t *testing.T) {
	expected := 26397
	// Given an input heightmap
	lines, _ := ParseInput("data/test-day-10.txt")

	// When the total corruption score is calculated
	score := Part1(lines)

	// The the correct response is returned
	util.AssertAnswer(expected, score, t)
}

func TestDay10Part2(t *testing.T) {
	expected := 288957
	// Given an input heightmap
	lines, _ := ParseInput("data/test-day-10.txt")

	// When the total completion score is calculated
	score := Part2(lines)

	// The the correct response is returned
	util.AssertAnswer(expected, score, t)
}
