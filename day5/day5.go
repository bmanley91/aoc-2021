package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"manley.dev/aoc/util"
)

type line struct {
	startPoint coordinate
	endPoint   coordinate
	direction  Direction
}

type coordinate struct {
	xVal int
	yVal int
}

type Direction int

const (
	Vertical   Direction = 0
	Horizontal Direction = 1
	Diagonal   Direction = 3
)

func Solve() {
	lines, err := ParseInputIntoLines("day5/data/day-5.txt")
	if err != nil {
		log.Fatalf("Couldn't read file %v", err)
	}

	fmt.Println("Day 5 Part 1: ", Part1(lines))
	fmt.Println("Day 5 Part 2: ", Part2(lines))
}

func ParseInputIntoLines(path string) (lines []line, err error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if err != nil {
		return nil, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()

		// Get Coordinates
		unformattedCoords := strings.Split(lineText, " -> ")
		coords := make([]coordinate, 2)
		for index, ununformattedCoordinate := range unformattedCoords {
			splitString := strings.Split(ununformattedCoordinate, ",")
			coords[index] = coordinate{
				xVal: util.StringToInt(splitString[0]),
				yVal: util.StringToInt(splitString[1]),
			}
		}

		// Determine direction
		if coords[0].xVal == coords[1].xVal {
			lines = append(lines, line{
				startPoint: coords[0],
				endPoint:   coords[1],
				direction:  Horizontal,
			})
		} else if coords[0].yVal == coords[1].yVal {
			lines = append(lines, line{
				startPoint: coords[0],
				endPoint:   coords[1],
				direction:  Vertical,
			})
		} else if math.Abs(float64(coords[0].yVal)-float64(coords[1].yVal)) ==
			math.Abs(float64(coords[0].xVal)-float64(coords[1].xVal)) {
			lines = append(lines, line{
				startPoint: coords[0],
				endPoint:   coords[1],
				direction:  Diagonal,
			})
		}
	}

	return lines, nil
}

func Part1(lines []line) (overlaps int) {
	// Create a grid
	var grid [1000][1000]int

	// Go through lines and mark the grid
	// Increment overlap count every time we hit 2 for a given coordinate
	for _, line := range lines {
		switch line.direction {
		case Vertical:
			min := int(math.Min(float64(line.startPoint.xVal), float64(line.endPoint.xVal)))
			max := int(math.Max(float64(line.startPoint.xVal), float64(line.endPoint.xVal)))
			for xCoord := min; xCoord <= max; xCoord++ {
				grid[line.startPoint.yVal][xCoord] += 1
				if grid[line.startPoint.yVal][xCoord] == 2 {
					overlaps++
				}
			}
		case Horizontal:
			min := int(math.Min(float64(line.startPoint.yVal), float64(line.endPoint.yVal)))
			max := int(math.Max(float64(line.startPoint.yVal), float64(line.endPoint.yVal)))
			for yCoord := min; yCoord <= max; yCoord++ {
				grid[yCoord][line.startPoint.xVal] += 1
				if grid[yCoord][line.startPoint.xVal] == 2 {
					overlaps++
				}
			}
		case Diagonal:
			continue
		default:
			log.Fatalf("Invalid line %v", line)
		}
	}

	return overlaps
}

func Part2(lines []line) (overlaps int) {
	// Create a grid
	var grid [1000][1000]int

	// Go through lines and mark the grid
	// Increment overlap count every time we hit 2 for a given coordinate
	for _, line := range lines {
		switch line.direction {
		case Vertical:
			min := int(math.Min(float64(line.startPoint.xVal), float64(line.endPoint.xVal)))
			max := int(math.Max(float64(line.startPoint.xVal), float64(line.endPoint.xVal)))
			for xCoord := min; xCoord <= max; xCoord++ {
				grid[line.startPoint.yVal][xCoord] += 1
				if grid[line.startPoint.yVal][xCoord] == 2 {
					overlaps++
				}
			}
		case Horizontal:
			min := int(math.Min(float64(line.startPoint.yVal), float64(line.endPoint.yVal)))
			max := int(math.Max(float64(line.startPoint.yVal), float64(line.endPoint.yVal)))
			for yCoord := min; yCoord <= max; yCoord++ {
				grid[yCoord][line.startPoint.xVal] += 1
				if grid[yCoord][line.startPoint.xVal] == 2 {
					overlaps++
				}
			}
		case Diagonal:
			// Go from top to bottom and figure out if we have to go left to right
			var topCoord, bottomCoord coordinate

			if line.startPoint.yVal < line.endPoint.yVal {
				topCoord = line.startPoint
				bottomCoord = line.endPoint
			} else {
				topCoord = line.endPoint
				bottomCoord = line.startPoint
			}

			totalRange := int(math.Abs(float64(topCoord.yVal) - float64(bottomCoord.yVal)))
			offset := 0

			if topCoord.xVal < bottomCoord.xVal {
				// Left to Right case
				for offset <= totalRange {
					grid[topCoord.yVal+offset][topCoord.xVal+offset] += 1
					if grid[topCoord.yVal+offset][topCoord.xVal+offset] == 2 {
						overlaps++
					}
					offset++
				}
			} else {
				// Right to Left case
				for offset <= totalRange {
					grid[topCoord.yVal+offset][topCoord.xVal-offset] += 1
					if grid[topCoord.yVal+offset][topCoord.xVal-offset] == 2 {
						overlaps++
					}
					offset++
				}
			}

		default:
			log.Fatalf("Invalid line %v", line)
		}
	}

	return overlaps
}
