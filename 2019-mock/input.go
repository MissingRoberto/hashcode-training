package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) Problem {
	file, err := os.Open(filePath)
	handleError("open file failed", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	params := readLineWithParams(scanner, 4, " ")
	problem := Problem{
		NumRows:    atoi(params[0]),
		NumColumns: atoi(params[1]),
		L:          atoi(params[2]),
		H:          atoi(params[3]),
	}

	grid := make([][]string, problem.NumRows)
	for index := 0; index < problem.NumRows; index++ {
		grid[index] = make([]string, problem.NumColumns)
		params := readLineWithParams(scanner, problem.NumColumns, "")
		for column := 0; column < problem.NumColumns; column++ {
			grid[index][column] = params[column]
		}
	}

	problem.Grid = grid
	err = scanner.Err()
	handleError("scanner failed", err)
	return problem
}

func readLineWithParams(scanner *bufio.Scanner, expected int, separator string) []string {
	params := strings.Split(readLine(scanner), separator)

	if len(params) != expected {
		log.Fatalf("first line must contain %d params", expected)
	}
	return params
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		log.Fatal("first line cannot be empty")
	}

	return scanner.Text()
}

func handleError(s string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", s, err)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	handleError("convert string to int failed", err)
	return i
}
