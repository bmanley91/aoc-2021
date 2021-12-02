package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadFileForInts(path string) []int {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Uh oh!", err)
	}

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Uh oh!", err)
		}

		lines = append(lines, intVal)
	}

	return lines
}

func ReadFileForStrings(path string) []string {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Uh oh!", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err != nil {
			log.Fatal("Uh oh!", err)
		}

		lines = append(lines, scanner.Text())
	}

	return lines
}
