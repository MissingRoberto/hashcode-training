package main

import "fmt"

func resolve(problem Problem) Result {
	return Result{}
}

func score(problem Problem, result Result) int {
	var score int
	return score
}

func printSolution(result Result) {
	fmt.Println(len(result))

	for _, s := range result {
		fmt.Printf("%d %d %d %d\n", s.row1, s.row2, s.column1, s.column2)
	}
}
