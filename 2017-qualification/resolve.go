package main

func resolve(problem Problem) Result {
	return resolveStupidToTwo(problem)
}

func resolveStupidToTwo(problem Problem) Result {
	for videoID, size := range problem.VideoSizes {
		counter := 2
		for i := 0; i < len(problem.Caches) && counter > 0; i++ {
			if problem.Caches[i].RemainingCapacity-size > 0 {
				problem.Caches[i].RemainingCapacity -= size
				problem.Caches[i].Videos[videoID] = true
				counter--
			}
		}
	}
	return problem.Caches
}

func score(problem Problem, result Result) int {
	var score int
	var totalRequests int
	for _, request := range problem.Requests {
		endpoint := problem.Endpoints[request.EndpointID]
		score += (endpoint.LatencyToDatacenter - chooseBest(request.VideoID, endpoint, result)) * request.Num
		totalRequests += request.Num
	}
	return score * 1000 / totalRequests
}

func chooseBest(video int, endpoint Endpoint, caches []Cache) int {
	best := endpoint.LatencyToDatacenter

	for id, latency := range endpoint.CacheLatencies {
		if caches[id].Videos[video] && latency < best {
			best = latency
		}
	}
	return best
}
