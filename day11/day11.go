package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"manley.dev/aoc/util"
)

type Offset struct {
	row int
	col int
}

var offsets = []Offset{
	{-1, -1}, // NW
	{-1, 0},  // N
	{-1, 1},  // NE
	{0, -1},  // W
	{0, 1},   // E
	{1, -1},  // SW
	{1, 0},   // S
	{1, 1},   // SE
}

const rounds = 100

var flashCount = 0

func Solve() {
	board, err := ParseInput("day11/data/day-11.txt")
	if err != nil {
		log.Fatalf("Couldn't load file %v", err)
	}

	// Since we're modifying the board, only one of these can actually be done at a time.
	fmt.Printf("Day 11 Part 1: %d\n", Part1(board))
	fmt.Printf("Day 11 Part 2: %d\n", Part2(board))
}

func ParseInput(path string) (board [][]int, err error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentLine []int
		text := scanner.Text()

		for _, numString := range strings.Split(text, "") {
			currentLine = append(currentLine, util.StringToInt(numString))
		}
		board = append(board, currentLine)
	}

	return board, nil
}

func Part1(energyLevels [][]int) int {
	// printBoard(energyLevels)
	isFlashing := setupIsFlashing(energyLevels)

	for i := 0; i < rounds; i++ {
		for row := 0; row < len(energyLevels); row++ {
			for col := 0; col < len(energyLevels); col++ {
				increaseEnergyLevel(energyLevels, isFlashing, row, col)
			}
		}
		// At the end of each round, we have to reset the flashing flags
		resetFlashingFlags(isFlashing)
	}

	// printBoard(energyLevels)

	return flashCount
}

func Part2(energyLevels [][]int) (roundCount int) {
	// Basically the same as part one, but we'll check if the flashing flags are all set to true at any given point
	// We'll loop until that happens
	isFlashing := setupIsFlashing(energyLevels)
	for !allAreFlashing(isFlashing) {
		roundCount++
		for row := 0; row < len(energyLevels); row++ {
			for col := 0; col < len(energyLevels); col++ {
				increaseEnergyLevel(energyLevels, isFlashing, row, col)
			}
		}

		// fmt.Printf("Round %d\n", roundCount)
		// printBoard(energyLevels)

		// Check if all are flashing before resetting
		if allAreFlashing(isFlashing) {
			return roundCount
		}

		// At the end of each round, we have to reset the flashing flags
		resetFlashingFlags(isFlashing)
	}

	return roundCount
}

func increaseEnergyLevel(energyLevels [][]int, isFlashing [][]bool, row int, col int) {

	// Base case, if we're out of bounds, don't try to do anything.
	if row < 0 || row >= len(energyLevels) || col < 0 || col >= len(energyLevels[0]) {
		return
	}

	// If this octopus is already marked as flashing,
	// everything should already be handled for it
	if isFlashing[row][col] {
		return
	}

	// If we're in bounds, and if the octopus isn't already flashing,
	// Then we increase its energy level
	newEnergyLevel := energyLevels[row][col] + 1

	// If the new energy level is 10,
	// // set its flashing flag to true,
	// // set its energy level to zero,
	// // increment the flash count
	// // and start recursing to other nodes
	if newEnergyLevel == 10 {
		newEnergyLevel = 0
		energyLevels[row][col] = newEnergyLevel
		isFlashing[row][col] = true
		flashCount++

		// Recurse in all directions
		for _, offset := range offsets {
			increaseEnergyLevel(
				energyLevels,
				isFlashing,
				row+offset.row,
				col+offset.col,
			)
		}
	} else {
		// If the energy level didn't hit 10, just set the energy level and we're done
		energyLevels[row][col] = newEnergyLevel
	}
}

func printBoard(board [][]int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			fmt.Printf("%d", board[row][col])
		}
		fmt.Println()
	}
	fmt.Println()
}

func setupIsFlashing(board [][]int) (isFlashing [][]bool) {
	for _, row := range board {
		isFlashing = append(isFlashing, make([]bool, len(row)))
	}
	return isFlashing
}

func resetFlashingFlags(input [][]bool) {
	for row := range input {
		for col := range input[row] {
			input[row][col] = false
		}
	}
}

func allAreFlashing(input [][]bool) bool {
	for row := range input {
		for col := range input[row] {
			if !input[row][col] {
				return false
			}
		}
	}
	return true
}
