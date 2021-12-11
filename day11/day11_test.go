package day11

import (
	"testing"

	"manley.dev/aoc/util"
)

func TestDay11Part1(t *testing.T) {
	expected := 1656
	// Given an input grid
	lines, _ := ParseInput("data/test-day-11.txt")

	// When the total number of flashes is calculated
	flashes := Part1(lines)

	// The the correct response is returned
	util.AssertAnswer(expected, flashes, t)
}

func TestDay11Part2(t *testing.T) {
	expected := 195
	// Given an input grid
	lines, _ := ParseInput("data/test-day-11.txt")

	// When the synchronization round is found
	rounds := Part2(lines)

	// The the correct response is returned
	util.AssertAnswer(expected, rounds, t)
}
