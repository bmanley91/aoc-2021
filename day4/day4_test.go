package day4

import (
	"testing"
)

func TestDay4Part1(t *testing.T) {
	const expected = 4512

	// Given a list of draws and some boards
	draws, boards := ProcessInput("data/test-day-4.txt")

	// When the winning score is calculated
	answer := Part1(draws, boards)

	// Then it matches the expected value
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}

func TestDay4Part2(t *testing.T) {
	const expected = 1924

	// Given a list of draws and some boards
	draws, boards := ProcessInput("data/test-day-4.txt")

	// When the winning score is calculated
	answer := Part2(draws, boards)

	// Then it matches the expected value
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}
