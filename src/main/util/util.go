package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileForInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		lines = append(lines, intVal)
	}

	return lines, scanner.Err()
}

func ReadFileForStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err != nil {
			fmt.Println(err)
		}

		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
