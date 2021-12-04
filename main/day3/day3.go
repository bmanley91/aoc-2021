package day3

import (
	"fmt"
	"math"

	"manley.dev/aoc/main/util"
)

func Solve() {
	readings := util.ReadFileForStrings("day3/data/day-3.txt")
	fmt.Printf("Day 3 Part 1: %d\n", Part1(readings))
}

func Part1(input []string) int {
	// Loop over each input and sum up the values at each position
	// Then divide by the number of inputs
	// If the result is > .5, then 1 is the most common
	// Otherwise 0 is the most common
	// *there's nothing in the problem about what to do if there's a tie*
	sums := make([]int, len(input[0]))

	for _, reading := range input {
		for index, char := range reading {
			if char == '1' {
				sums[index]++
			}
		}
	}

	gamma, epsilon := findGammaAndEpsilon(sums, len(input))

	return gamma * epsilon
}

// Gamma binary value is the most common bits

// Epsilon is the inverse of Gamma
// (e.g. 101 => 010)
func findGammaAndEpsilon(sumArray []int, entryCount int) (gamma int, epsilon int) {
	mostCommon := make([]int, len(sumArray))
	leastCommon := make([]int, len(sumArray))

	for index, value := range sumArray {
		mostCommon[index] = int(math.Round(float64(value) / float64(entryCount)))
		if mostCommon[index] == 1 {
			leastCommon[index] = 0
		} else {
			leastCommon[index] = 1
		}
	}

	for i := 0; i < len(mostCommon); i++ {
		if mostCommon[i] == 1 {
			gamma += int(math.Pow(2, float64(len(mostCommon)-i-1)))
		}

		if leastCommon[i] == 1 {
			epsilon += int(math.Pow(2, float64(len(mostCommon)-i-1)))
		}
	}

	return gamma, epsilon
}
