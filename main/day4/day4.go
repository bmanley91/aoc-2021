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

type board struct {
	b [][]cell
}

func Solve() {
	fmt.Println("Day 4 Part 1", Part1())
}

func Part1() int {
	draws, boards := processInput("day4/data/day-4.txt")
	firstWinner, lastDraw := findFirstWinner(draws, boards)
	return firstWinner.score(lastDraw)
}

func processInput(path string) (draws []int, boards []board) {
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

func parseBoards(blocks []string) []board {
	boards := make([]board, len(blocks))
	for i, s := range blocks {
		board := parseBoard(s)
		boards[i] = board
	}
	return boards
}

func parseBoard(raw string) board {
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
	return board{b: b}
}

func findFirstWinner(draws []int, boards []board) (winner board, lastDraw int) {
	bs := make([]board, len(boards))
	copy(bs, boards)

	for _, currentDraw := range draws {
		for _, currentBoard := range bs {
			currentBoard.mark(currentDraw)
			if currentBoard.isWinner() {
				return currentBoard, currentDraw
			}
		}
	}
	return board{}, -1
}

func (board board) mark(val int) {
	for i := range board.b {
		for j := range board.b[i] {
			if board.b[i][j].num == val {
				board.b[i][j].called = true
			}
		}
	}
}

func (board board) isWinner() bool {
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

func (board board) score(lastDraw int) int {
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
