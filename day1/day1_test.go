package day1

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	expected := 7
	// Given a list of depth values
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	// When we measure the number of depth increases
	answer := Part1(input)

	// Then we get the correct answer
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}

func TestDay1Part2(t *testing.T) {
	expected := 5
	// Given a list of depth values
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	// When we measure the number of depth increases
	answer := Part2(input)

	// Then we get the correct answer
	if answer != expected {
		t.Errorf("Incorrect answer. Expected %d, got %d", expected, answer)
	}
}
