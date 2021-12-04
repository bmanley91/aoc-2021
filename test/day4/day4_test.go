package day4

import (
	"testing"

	"manley.dev/aoc/main/day4"
)

func TestDay4Part1(t *testing.T) {
	const expected = 4512

	// Given a list of draws and some boards
	draws, boards := day4.ProcessInput("data/test-day-4.txt")

	// When the winning score is calculated
	answer := day4.Part1(draws, boards)

	// Then it matches the expected value
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}

func TestDay4Part2(t *testing.T) {
	const expected = 1924

	// Given a list of draws and some boards
	draws, boards := day4.ProcessInput("data/test-day-4.txt")

	// When the winning score is calculated
	answer := day4.Part2(draws, boards)

	// Then it matches the expected value
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}
