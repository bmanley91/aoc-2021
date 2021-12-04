package day4

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Boards are always 5 rows long
const size = 5

type cell struct {
	num    int
	called bool
}

type Board struct {
	b [][]cell
}

func Solve() {
	draws, boards := ProcessInput("day4/data/day-4.txt")
	fmt.Println("Day 4 Part 1: ", Part1(draws, boards))
	fmt.Println("Day 4 Part 2: ", Part2(draws, boards))
}

func Part1(draws []int, boards []Board) int {
	firstWinner, lastDraw := findFirstWinner(draws, boards)
	return firstWinner.score(lastDraw)
}

func Part2(draws []int, boards []Board) int {
	lastWinner, lastDraw := findLastWinner(draws, boards)
	return lastWinner.score(lastDraw)
}

func ProcessInput(path string) (draws []int, boards []Board) {
	absPath, _ := filepath.Abs(path)
	file, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatal("Uh oh!", err)
	}

	blocks := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	// Draws will always be the first block
	draws = parseDraws(blocks[0])

	// Boards will be everything after the first block
	boards = parseBoards(blocks[1:])

	return draws, boards
}

func parseDraws(input string) []int {
	split := strings.Split(input, ",")
	draws := make([]int, len(split))

	for i, s := range split {
		d, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Uh oh! ", err)
		}
		draws[i] = int(d)
	}

	return draws
}

func parseBoards(blocks []string) []Board {
	boards := make([]Board, len(blocks))
	for i, s := range blocks {
		board := parseBoard(s)
		boards[i] = board
	}
	return boards
}

func parseBoard(raw string) Board {
	rows := strings.Split(raw, "\n")
	b := make([][]cell, size)
	for i := 0; i < size; i++ {
		b[i] = make([]cell, size)

		row := strings.Fields(strings.TrimSpace(rows[i]))
		for j := 0; j < size; j++ {
			num, err := strconv.Atoi(row[j])
			if err != nil {
				log.Fatal("Uh oh! ", err)
			}
			b[i][j] = cell{num: num}
		}
	}
	return Board{b: b}
}

func findFirstWinner(draws []int, boards []Board) (winner Board, lastDraw int) {
	bs := make([]Board, len(boards))
	copy(bs, boards)

	for _, currentDraw := range draws {
		for _, currentBoard := range bs {
			currentBoard.mark(currentDraw)
			if currentBoard.isWinner() {
				return currentBoard, currentDraw
			}
		}
	}

	// Default, should never happen?
	return Board{}, -1
}

func findLastWinner(draws []int, boards []Board) (winner Board, lastDraw int) {
	currentBoardArray := make([]Board, len(boards))
	copy(currentBoardArray, boards)

	for _, currentDraw := range draws {
		nextBoardArray := []Board{}
		for _, currentBoard := range currentBoardArray {
			currentBoard.mark(currentDraw)
			if len(currentBoardArray) == 1 && currentBoard.isWinner() {
				return currentBoard, currentDraw
			} else {
				if !currentBoard.isWinner() {
					nextBoardArray = append(nextBoardArray, currentBoard)
				}
			}
		}
		currentBoardArray = nextBoardArray
	}

	// Default, shouldn't happen?
	return Board{}, -1
}

func (board Board) mark(val int) {
	for i := range board.b {
		for j := range board.b[i] {
			if board.b[i][j].num == val {
				board.b[i][j].called = true
			}
		}
	}
}

func (board Board) isWinner() bool {
	// Check rows
	for i := 0; i < size; i++ {
		w := true
		for j := 0; j < size; j++ {
			w = w && board.b[i][j].called
		}
		if w {
			return true
		}
	}
	// Check columns
	for j := 0; j < size; j++ {
		w := true
		for i := 0; i < size; i++ {
			w = w && board.b[i][j].called
		}
		if w {
			return true
		}
	}
	return false
}

func (board Board) score(lastDraw int) int {
	var s int
	for i := range board.b {
		for j := range board.b[i] {
			if !board.b[i][j].called {
				s += int(board.b[i][j].num)
			}
		}
	}
	return s * lastDraw
}
