package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"manley.dev/aoc/main/util"
)

const initialDaysToSpawn = 6
const newFish = 8
const newFishExtraDaysToSpawn = 2

func Solve() {
	fishList, err := ParseInputIntoFish("day6/data/day-6.txt")

	if err != nil {
		log.Fatalf("Couldn't read file %v", err)
	}

	fmt.Printf("Day 6 Part 1: %d\n", Part1(fishList))
	fmt.Printf("Day 6 Part 2: %d\n", Part2(fishList))
}

func ParseInputIntoFish(path string) (fishList []int, err error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read text
		text := scanner.Text()
		splitString := strings.Split(text, ",")

		// Convert values to fish
		for _, fishValue := range splitString {
			fishList = append(fishList, util.StringToInt(fishValue))
		}
	}

	// fmt.Printf("Initial fish list %v\n", fishList)
	return fishList, nil
}

func Part1(fishList []int) int {
	return findFishCount(fishList, 80)
}

func Part2(fishList []int) int {
	return findFishCountBetter(fishList, 256)
}

func findFishCount(fishList []int, days int) int {
	for currentDay := 0; currentDay < days; currentDay++ {
		var newFishList []int
		// Iterate over each fish
		for index, currentFish := range fishList {
			if currentFish == 0 {
				// Spawn Case
				// Create new fish and add it to the new fish list
				newFish := initialDaysToSpawn + newFishExtraDaysToSpawn
				newFishList = append(newFishList, newFish)

				// Reset current fish's days until next spawn
				fishList[index] = initialDaysToSpawn
			} else {
				// Non-Spawn Case
				// Decrement days until next spawn
				// fmt.Printf("Decrementing fish %v\n", currentFish)
				currentFish -= 1
				fishList[index] = currentFish
				// fmt.Printf("Updated fish %v\n", currentFish)
			}
		}

		fishList = append(fishList, newFishList...)
		// fmt.Printf("Current day: %v, Current count: %v\n", currentDay, len(fishList))
	}

	return len(fishList)
}

func findFishCountBetter(fishList []int, days int) (sum int) {
	counter := initializeCounter(fishList)

	for currentDay := 0; currentDay < days; currentDay++ {
		nextCounter := make(map[int]int)
		for numDays, count := range counter {
			if numDays == 0 {
				// Move all of the fish that are due to spawn to 6
				nextCounter[initialDaysToSpawn] += count

				// Add all of the new fish to 8
				nextCounter[newFish] += count
			} else {
				// Move all of the fish from the current day down one spot
				// (e.g.) 10 fish at day 2 go to day 1
				nextCounter[numDays-1] += count
			}
		}
		counter = nextCounter
	}

	fmt.Printf("counter: %v\n", counter)

	// Iterate through the map and sum up all the values
	for _, value := range counter {
		sum += value
	}

	return sum
}

func initializeCounter(fishList []int) map[int]int {
	counterMap := make(map[int]int)

	for _, fish := range fishList {
		if _, present := counterMap[fish]; present {
			counterMap[fish]++
		} else {
			counterMap[fish] = 1
		}
	}

	return counterMap
}

func sum(counter map[int]int) int {
	count := 0
	for _, val := range counter {
		count += val
	}
	return count
}
