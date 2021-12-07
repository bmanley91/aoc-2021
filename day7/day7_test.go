package day7

import (
	"testing"

	"manley.dev/aoc/util"
)

func TestDay7Part1(t *testing.T) {
	expected := 37
	// Given an input list of positions
	positions, _ := ParseInput("data/test-day-7.txt")

	// When the minimum fuel cost is calculated
	answer := Part1(positions)

	// The the correct response is returned
	util.AssertAnswer(expected, answer, t)
}
func TestDay7Part2(t *testing.T) {
	expected := 168
	// Given an input list of positions
	positions, _ := ParseInput("data/test-day-7.txt")

	// When the minimum fuel cost is calculated
	answer := Part2(positions)

	// The the correct response is returned
	util.AssertAnswer(expected, answer, t)
}
