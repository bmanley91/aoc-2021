package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"manley.dev/aoc/util"
)

func Solve() {
	heights, err := ParseInput("day9/data/day-9.txt")

	if err != nil {
		log.Fatal("Failed to read file %w", err)
	}

	fmt.Printf("Day 9 Part 1: %d\n", Part1(heights))
	fmt.Printf("Day 9 Part 2: %d\n", Part2(heights))
}

func ParseInput(path string) ([][]int, error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	var heights [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		heightChars := strings.Split(lineText, "")

		var heightsForLine []int
		for _, value := range heightChars {
			heightsForLine = append(heightsForLine, util.StringToInt(value))
		}

		heights = append(heights, heightsForLine)
	}

	return heights, nil
}

func Part1(heights [][]int) (riskLevelSum int) {
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[0]); col++ {
			if isLowPoint(heights, row, col) {
				riskLevelSum += calculateRiskScore(heights, row, col)
			}
		}
	}
	return riskLevelSum
}

func Part2(heights [][]int) int {
	// This is the island area problem except we need to check
	// adjacent point values <= current point value
	// Also we need to keep track of the top three sizes
	return -1
}

func isLowPoint(heights [][]int, row int, col int) bool {
	cellValue := heights[row][col]
	northIsHigher := (row-1 >= 0 &&
		cellValue < heights[row-1][col]) ||
		row-1 < 0
	southIsHigher := (row+1 < len(heights) &&
		cellValue < heights[row+1][col]) ||
		row+1 >= len(heights)
	westIsHigher := (col-1 >= 0 &&
		cellValue < heights[row][col-1]) ||
		col-1 < 0
	eastIsHigher := (col+1 < len(heights[0]) &&
		cellValue < heights[row][col+1]) ||
		col+1 >= len(heights[0])

	return northIsHigher && southIsHigher && westIsHigher && eastIsHigher
}

func calculateRiskScore(heights [][]int, row int, col int) int {
	// fmt.Printf(
	// 	"Low point row: %d, col: %d has a value of %d and a risk of %d\n",
	// 	row,
	// 	col,
	// 	heights[row][col],
	// 	heights[row][col]+1,
	// )
	return heights[row][col] + 1
}
