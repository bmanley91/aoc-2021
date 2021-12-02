package day1

import (
	"fmt"

	"manley.dev/aoc/src/main/util"
)

func Solve() {
	// Read data into array
	depthValues := util.ReadFileForInts("src/main/day1/data/day-1.txt")

	part1(depthValues)
	part2(depthValues)
}

func part1(depthValues []int) {
	// Loop through list with two pointers
	// If value at previous pointer is less than value at current pointer,
	// increment a running total of depth increases
	depthIncreaseCount := 0
	previous := 0
	current := 1

	for current < len(depthValues) {
		if depthValues[previous] < depthValues[current] {
			depthIncreaseCount++
		}

		previous++
		current++
	}

	// Return running total once we've read all lines
	fmt.Printf("Number of depth increases: %d\n", depthIncreaseCount)
}

func part2(depthValues []int) {
	// Each window will share two values so we actually only care about the values around the two shared ones
	// i.e. looking at the windows [0,1,2] and [1,2,3], we need to compare the values at positions 0 and 3
	depthWindowIncreases := 0
	previous := 0
	current := previous + 3

	for current < len(depthValues) {
		if depthValues[previous] < depthValues[current] {
			depthWindowIncreases++
		}

		previous++
		current++
	}

	fmt.Printf("Number of depth increases: %d\n", depthWindowIncreases)
}
