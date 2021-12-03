package day3

import (
	"testing"

	"manley.dev/aoc/main/day3"
)

func TestDay3Part1(t *testing.T) {
	expected := 198
	// Given an input list
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	// When we calculate the answer
	answer := day3.Part1(input)

	// Then it matches the expected value
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}
