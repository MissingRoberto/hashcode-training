package main

import (
	"math"
	"sort"
)

func Dist(a, b Position) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func ScoreRide(ride Ride, bonus int, t int, idleCars []Car) int {
	dist := Dist(ride.From, ride.To)
	if dist+t > ride.LatestFinish {
		return 0
	}
	score := dist
	if ride.EarliestStart <= t {
		score += bonus - (t - ride.EarliestStart)
	}
	// + Probability of Bonus * Bonus
	return score
}

func Score(car Car, ride Ride, bonus int, t int) int {
	distFromCar := Dist(car.Position, ride.From)
	dist := Dist(ride.From, ride.To)
	if dist+distFromCar+t > ride.LatestFinish {
		return 0
	}
	score := dist

	if ride.EarliestStart <= distFromCar+t {
		score += bonus
	}
	return score
}

func GreedyScore(car Car, ride Ride, bonus, t int) int {
	score := Score(car, ride, bonus, t)
	return score - WaitTime(car, ride, t) - Dist(car.Position, ride.From)
}

func PickupTime(car Car, ride Ride, t int) int {
	return t + Dist(car.Position, ride.From)
}

func WaitTime(car Car, ride Ride, t int) int {
	pickupTime := t + Dist(car.Position, ride.From)
	if pickupTime >= ride.EarliestStart {
		return 0
	}
	return ride.EarliestStart - pickupTime
}

func DropTime(car Car, ride Ride, t int) int {
	return t + Dist(car.Position, ride.From) + Dist(ride.From, ride.To) + WaitTime(car, ride, t)
}

func Resolve(problem Problem) int {
	idleCars := problem.Cars
	remainingRides := problem.Rides
	arrivals := make([][]Car, problem.Steps)

	var score int

	for t := 0; t < problem.Steps; t++ {
		idleCars = append(idleCars, arrivals[t]...)
		if len(idleCars) == 0 {
			continue
		}

		for len(idleCars) > 0 && len(remainingRides) > 0 {
			// Pick a ride
			sort.Slice(remainingRides, func(i, j int) bool {
				return ScoreRide(remainingRides[i], problem.Bonus, t, idleCars) < ScoreRide(remainingRides[j], problem.Bonus, t, idleCars)
			})

			ride := remainingRides[len(remainingRides)-1]
			remainingRides = remainingRides[:len(remainingRides)-1]

			sort.Slice(idleCars, func(i, j int) bool {
				return GreedyScore(idleCars[i], ride, problem.Bonus, t) < GreedyScore(idleCars[i], ride, problem.Bonus, t)
			})

			car := idleCars[len(idleCars)-1]
			idleCars = idleCars[:len(idleCars)-1]

			score += Score(car, ride, problem.Bonus, t)
			dropTime := DropTime(car, ride, t)
			if dropTime < problem.Steps {
				arrivals[dropTime] = append(arrivals[dropTime], Car{Position: ride.To})
			}
		}
	}
	return score
}

func ResolveBasicGreedy(problem Problem) int {
	idleCars := problem.Cars
	remainingRides := problem.Rides
	arrivals := make([][]Car, problem.Steps)

	var score int

	for t := 0; t < problem.Steps; t++ {
		idleCars = append(idleCars, arrivals[t]...)
		if len(idleCars) == 0 {
			continue
		}

		previousCar := idleCars[0]

		for len(idleCars) > 0 && len(remainingRides) > 0 {
			// Pick a car
			sort.Slice(idleCars, func(i, j int) bool {
				return Dist(idleCars[i].Position, previousCar.Position) > Dist(idleCars[j].Position, previousCar.Position)
			})

			car := idleCars[len(idleCars)-1]
			idleCars = idleCars[:len(idleCars)-1]

			// Pick a ride
			sort.Slice(remainingRides, func(i, j int) bool {
				return Score(car, remainingRides[i], problem.Bonus, t) < Score(car, remainingRides[j], problem.Bonus, t)
			})

			ride := remainingRides[len(remainingRides)-1]
			remainingRides = remainingRides[:len(remainingRides)-1]

			score += Score(car, ride, problem.Bonus, t)
			dropTime := DropTime(car, ride, t)
			if dropTime < problem.Steps {
				arrivals[dropTime] = append(arrivals[dropTime], Car{Position: ride.To})
			}
		}
	}
	return score
}
