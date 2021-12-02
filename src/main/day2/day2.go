package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"manley.dev/aoc/src/main/util"
)

const directionIndex = 0
const amountIndex = 1

func Solve() {
	// Read data into array
	instructions, readError := util.ReadFileForStrings("main/day2/data/day-2.txt")

	if readError != nil {
		log.Fatalf("Uh oh!", readError)
	}

	fmt.Printf("Part 1: %d\n", Part1(instructions))
}

func Part1(instructions []string) int {
	depth, horizontal := followInstructions(instructions)
	return determineAnswer(depth, horizontal)
}

func followInstructions(instructions []string) (depth int, horizontal int) {
	for _, instruction := range instructions {
		splitInstruction := strings.Split(instruction, " ")

		switch splitInstruction[directionIndex] {
		case "forward":
			horizontal += getAmountFromInstruction(splitInstruction)

		case "up":
			depth -= getAmountFromInstruction(splitInstruction)

		case "down":
			depth += getAmountFromInstruction(splitInstruction)

		default:
			fmt.Printf("Uhhh what? %s", instruction)
		}
	}

	return depth, horizontal
}

func getAmountFromInstruction(instruction []string) int {
	amount, _ := strconv.Atoi(instruction[amountIndex])
	return amount
}

func determineAnswer(depth int, horizontal int) int {
	return depth * horizontal
}