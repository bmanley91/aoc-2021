package day6

import "fmt"

type Fish struct {
	daysUntilNextSpawn int
	generation         int
}

func Solve() {
	fishList := ParseInputIntoFish("day6/data/day-6.txt")

	fmt.Printf("Day 6 Part 1: %d", Part1(fishList))
}

func ParseInputIntoFish(path string) (fishList []Fish) {
	return fishList
}

func Part1(fishList []Fish) int {
	return -1
}
