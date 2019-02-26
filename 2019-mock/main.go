package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var totalScore int
	for i := 0; i < len(args); i++ {
		problem := readInput(args[i])
		fmt.Printf("%+v\n", problem)
		result := resolve(problem)
		printSolution(result)
		// fmt.Printf("Result %s: %+v\n", args[i], result)
		score := score(problem, result)
		fmt.Printf("Score %s: %d\n", args[i], score)
		totalScore += score
	}
	fmt.Printf("Total Score: %d\n", totalScore)
}
