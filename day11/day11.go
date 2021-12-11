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

func Solve() {
	board, err := ParseInput("day11/data/day-11.txt")
	if err != nil {
		log.Fatalf("Couldn't load file %v", err)
	}

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

func Part1(board [][]int) int {
	return -1
}

func Part2(board [][]int) int {
	return -1
}
