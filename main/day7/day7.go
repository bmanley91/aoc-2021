package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"manley.dev/aoc/main/util"
)

func Solve() {
	positions, err := ParseInput("day7/data/day-7.txt")
	if err != nil {
		log.Fatal("Error reading file %w", err)
	}

	fmt.Printf("Day 7 Part 1: %v\n", Part1(positions))
	fmt.Printf("Day 7 Part 2: %v\n", Part2(positions))
}

func ParseInput(path string) ([]int, error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	var positions []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		positionStrings := strings.Split(lineText, ",")
		for _, text := range positionStrings {
			positions = append(positions, util.StringToInt(text))
		}
	}

	return positions, nil
}

func Part1(positions []int) int {
	// Find the median position
	median := findMedianPosition(positions)

	// Iterate through each position and sum up their distances from the median position
	// Return the sum
	return findTotalCostToMove(positions, median)
}

func Part2(positions []int) int {
	// Find the average position
	average := findAveragePosition(positions)
	return findCostToMoveWithMoreGas(positions, average)
}

func findMedianPosition(positions []int) int {
	sort.Ints(positions)

	middleIndex := len(positions) / 2

	if len(positions)%2 == 1 {
		return positions[middleIndex]
	}

	return (positions[middleIndex-1] + positions[middleIndex]) / 2
}

func findTotalCostToMove(positions []int, target int) (sum int) {
	for _, position := range positions {
		sum += int(math.Abs(float64(position) - float64(target)))
	}
	return sum
}

func findAveragePosition(positions []int) int {
	sum := 0
	for _, position := range positions {
		sum += position
	}
	return int(math.Round(float64(sum) / float64(len(positions))))
}

func findCostToMoveWithMoreGas(positions []int, target int) (sum int) {
	for _, position := range positions {
		start := int(math.Min(float64(position), float64(target)))
		end := int(math.Max(float64(position), float64(target)))

		gasForPosition := 0
		for offset := 0; start+offset <= end; offset++ {
			gasForPosition += offset
			// fmt.Printf("Start: %d, Offset: %d\n", start, offset)
		}
		// fmt.Printf("Gas for position %d to move to target %d is %d\n", position, target, gasForPosition)

		sum += gasForPosition
	}

	return sum
}
