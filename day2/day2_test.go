package day2

import (
	"testing"
)

func TestDay2Part1(t *testing.T) {
	expected := 150
	// Given a list of instructions
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	// When we follow the instructions
	answer := Part1(input)

	// Then we get the correct depth and height
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}

func TestDay2Part2(t *testing.T) {
	expected := 900
	// Given a list of instructions
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	// When we follow the instructions
	answer := Part2(input)

	// Then we get the correct depth and height
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}
