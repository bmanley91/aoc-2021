package day2

import (
	"fmt"
	"log"

	"manley.dev/aoc/src/main/util"
)

func Day2(path string) {
	// Read data into array
	instructions, readError := util.ReadFileForStrings(path)

	if readError != nil {
		log.Fatalf("Uh oh!")
	}

	fmt.Printf("Part 1: %d", Part1(instructions))
}

func Part1(instructions []string) int {
	depth, horizontal := followInstructions(instructions)
	return determineAnswer(depth, horizontal)
}

func followInstructions(instructions []string) (depth int, horizontal int) {
	return
}

func determineAnswer(depth int, horizontal int) int {
	return depth * horizontal
}
