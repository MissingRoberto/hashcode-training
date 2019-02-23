package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Problem defines the input parameters
type Problem struct {
	Rows     int
	Columns  int
	NumCars  int
	NumRides int
	Bonus    int
	Steps    int
	Rides    []Ride
	Cars     []Car
}

type Ride struct {
	From          Position
	To            Position
	EarliestStart int
	LatestFinish  int
}

type Car struct {
	Position Position
}

type Position struct {
	x, y int
}

func main() {
	args := os.Args[1:]
	var totalScore int
	for i := 0; i < len(args); i++ {
		score := resolve(args[i])
		fmt.Printf("Score %s: %d\n", args[i], score)
		totalScore += score
	}
	fmt.Printf("Total Score: %d\n", totalScore)
}

func resolve(filePath string) int {
	file, err := os.Open(filePath)
	handleError("open file failed", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("first line cannot be empty")
	}

	params := strings.Split(scanner.Text(), " ")

	if len(params) != 6 {
		log.Fatal("first line must contain 6 params")
	}

	problem := Problem{
		Rows:     atoi(params[0]),
		Columns:  atoi(params[1]),
		NumCars:  atoi(params[3]),
		NumRides: atoi(params[3]),
		Bonus:    atoi(params[4]),
		Steps:    atoi(params[5]),
	}

	rides := make([]Ride, problem.NumRides)

	for index := 0; index < problem.NumRides; index++ {
		if !scanner.Scan() {
			log.Fatal("ride line cannot be empty")
		}

		params := strings.Split(scanner.Text(), " ")
		if len(params) != 6 {
			log.Fatal("each ride must contain 6 params")
		}

		rides = append(rides, Ride{
			From: Position{
				x: atoi(params[0]),
				y: atoi(params[1]),
			},
			To: Position{
				x: atoi(params[2]),
				y: atoi(params[3]),
			},
			EarliestStart: atoi(params[4]),
			LatestFinish:  atoi(params[5]),
		})
	}

	problem.Rides = rides
	problem.Cars = make([]Car, problem.NumCars)

	err = scanner.Err()
	handleError("scanner failed", err)

	// fmt.Printf("Problem: %+v\n", problem)
	// fmt.Printf("Rides: %+v\n", rides)
	return Resolve(problem)
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
