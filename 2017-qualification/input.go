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

	params := readLineWithParams(scanner, 5)
	problem := Problem{
		NumVideos:        atoi(params[0]),
		NumEndpoints:     atoi(params[1]),
		NumRequestBursts: atoi(params[2]),
		NumCaches:        atoi(params[3]),
		MaxCapacity:      atoi(params[4]),
	}

	caches := make([]Cache, problem.NumCaches)
	for i := range caches {
		caches[i] = Cache{
			RemainingCapacity: problem.MaxCapacity,
			Videos:            make(map[int]bool),
		}
	}
	problem.Caches = caches

	videos := make([]int, problem.NumVideos)
	params = readLineWithParams(scanner, problem.NumVideos)
	for i, p := range params {
		videos[i] = atoi(p)
	}

	problem.VideoSizes = videos

	endpoints := make([]Endpoint, problem.NumEndpoints)
	for index := 0; index < problem.NumEndpoints; index++ {
		params = readLineWithParams(scanner, 2)
		latencyToDatacenter, numCacheConnections := atoi(params[0]), atoi(params[1])

		latencies := make(map[int]int)
		for j := 0; j < numCacheConnections; j++ {
			params = readLineWithParams(scanner, 2)
			latencies[atoi(params[0])] = atoi(params[1])
		}
		endpoints[index] = Endpoint{
			LatencyToDatacenter: latencyToDatacenter,
			NumCaches:           numCacheConnections,
			CacheLatencies:      latencies,
		}
	}
	problem.Endpoints = endpoints
	requestBursts := make([]RequestBurst, problem.NumRequestBursts)

	for index := 0; index < problem.NumRequestBursts; index++ {
		params = readLineWithParams(scanner, 3)

		requestBursts[index] = RequestBurst{
			VideoID:    atoi(params[0]),
			EndpointID: atoi(params[1]),
			Num:        atoi(params[2]),
		}
	}
	problem.Requests = requestBursts

	err = scanner.Err()
	handleError("scanner failed", err)
	return problem
}

func readLineWithParams(scanner *bufio.Scanner, expected int) []string {
	params := strings.Split(readLine(scanner), " ")

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
