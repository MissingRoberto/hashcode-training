package main

import "math"

func Dist(a, b Position) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func Score(car Car, ride Ride, bonus int, t int) int {
	distFromCar := Dist(car.Position, ride.From)
	dist := Dist(ride.From, ride.To)
	if dist+distFromCar+t > ride.LatestFinish {
		return 0
	}
	score := dist

	if ride.EarliestStart == t {
		score += bonus
	}
	return score
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

		for len(idleCars) > 0 && len(remainingRides) > 0 {
			// Pick a car
			car := idleCars[len(idleCars)-1]
			idleCars = idleCars[:len(idleCars)-1]
			// Pick a ride
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
