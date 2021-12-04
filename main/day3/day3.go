package day3

import (
	"fmt"
	"math"

	"manley.dev/aoc/main/util"
)

func Solve() {
	readings := util.ReadFileForStrings("day3/data/day-3.txt")
	fmt.Printf("Day 3 Part 1: %d\n", Part1(readings))
	fmt.Printf("Day 3 Part 2: %d\n", Part2(readings))
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

func Part2(input []string) int {
	// Find most common like in Part1
	sums := make([]int, len(input[0]))
	for _, reading := range input {
		for index, char := range reading {
			if char == '1' {
				sums[index]++
			}
		}
	}

	mostCommon := make([]int, len(sums))
	for index, value := range sums {
		mostCommon[index] = int(math.Round(float64(value) / float64(len(input))))
	}
	fmt.Println("Most common ", mostCommon)

	findOxygenRating(input, mostCommon)

	return 0
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

func findOxygenRating(inputData []string, mostCommonBits []int) {
	recordsToKeep := make([]string, len(inputData))
	for recordIndex, data := range inputData {
		shouldAdd := true
		for index, bit := range inputData[recordIndex] {
			if mostCommonBits[index] != int(bit-'0') {
				fmt.Printf("Kicking out %s because of position %d. %s does not match %s\n",
					data, index, mostCommonBits[index], int(bit-'0'))
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			recordsToKeep = append(recordsToKeep, data)
		}
	}

	fmt.Println("this? ", recordsToKeep)
}
