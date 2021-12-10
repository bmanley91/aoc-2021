package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return ' '
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

var closingBracketCorruptionScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var bracketIncompletenessScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var bracketMatches = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Solve() {
	inputLines, err := ParseInput("day10/data/day-10.txt")

	if err != nil {
		log.Fatal("Failed to read file %w", err)
	}

	fmt.Printf("Day 10 Part 1: %d\n", Part1(inputLines))
	fmt.Printf("Day 10 Part 2: %d\n", Part2(inputLines))
}

func ParseInput(path string) (output []string, err error) {
	absPath, _ := filepath.Abs(path)
	file, e := os.Open(absPath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output, nil
}

func Part1(inputLines []string) (sum int) {
	for _, line := range inputLines {
		// Iterate through each char in the string
		// If the char is an opening bracket, add it to the stack
		// It it's a closing bracket, pop the top element off of the stack and see if they match
		// // If they don't then the line is corrupt and we should add the score for the bad closing bracket to the running total
		sum += getCorruptionScore(line)
	}

	return sum
}

func Part2(inputLines []string) int {
	var incompletenessScores []int
	for _, line := range inputLines {
		// We only care about incomplete lines
		score := getIncompletenessScore(line)
		if score != 0 {
			incompletenessScores = append(incompletenessScores, score)
		}
	}

	// Find the middle entry
	sort.Ints(incompletenessScores)
	return incompletenessScores[len(incompletenessScores)/2]
}

func getCorruptionScore(line string) int {
	var stack Stack
	for _, currentChar := range line {

		if isClosingBracket(currentChar) {
			previousOpenBracket := stack.Pop()
			if !bracketsMatch(previousOpenBracket, currentChar) {
				return getCorruptionScoreForBracket(currentChar)
			}
		} else {
			stack.Push(currentChar)
		}
	}

	return 0
}

func getIncompletenessScore(line string) int {
	// Go through the line and do our stack operations
	// If there's anything left in the stack at the end, then the line is incomplete
	// We'll need to figure out what chars need to be added to empty the stack
	var stack Stack
	for _, currentChar := range line {
		if isClosingBracket(currentChar) {
			previousOpenBracket := stack.Pop()
			if !bracketsMatch(previousOpenBracket, currentChar) {
				return 0
			}
		} else {
			stack.Push(currentChar)
		}
	}

	if stack.IsEmpty() {
		return 0
	}

	return scoreToCompleteStack(stack)
}

func isClosingBracket(input rune) bool {
	if _, ok := closingBracketCorruptionScores[input]; ok {
		return true
	}
	return false
}

func getCorruptionScoreForBracket(input rune) int {
	return closingBracketCorruptionScores[input]
}

func bracketsMatch(open rune, close rune) bool {
	return close == bracketMatches[open]
}

func scoreToCompleteStack(stack Stack) (sum int) {
	for !stack.IsEmpty() {
		sum *= 5
		currentChar := stack.Pop()
		sum += getIncompletenessScoreForBracket(currentChar)
	}
	return sum
}

func getIncompletenessScoreForBracket(input rune) int {
	return bracketIncompletenessScores[input]
}
