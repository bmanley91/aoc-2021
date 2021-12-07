package day7

import (
	"testing"

	"manley.dev/aoc/main/day7"
	"manley.dev/aoc/main/util"
)

func TestDay7Part1(t *testing.T) {
	expected := 37
	// Given an input list of positions
	positions, _ := day7.ParseInput("data/test-day-7.txt")

	// When the minimum fuel cost is calculated
	answer := day7.Part1(positions)

	// The the correct response is returned
	util.AssertAnswer(expected, answer, t)
}
func TestDay7Part2(t *testing.T) {

}
