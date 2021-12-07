package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"manley.dev/aoc/main/util"
)

func Solve() {
	positions, err := ParseInput("day7/data/day-7.txt")
	if err != nil {
		log.Fatal("Error reading file %w", err)
	}

	fmt.Printf("Day 7 Part 1: %v\n", Part1(positions))
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
	return -1
}
