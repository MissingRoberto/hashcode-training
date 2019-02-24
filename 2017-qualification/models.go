package main

// Problem defines the input parameters
type Problem struct {
	NumVideos        int
	NumEndpoints     int
	NumRequestBursts int
	NumCaches        int
	MaxCapacity      int
	VideoSizes       []int
	Endpoints        []Endpoint
	Caches           []Cache
	Requests         []RequestBurst
}

type Result []Cache

type Endpoint struct {
	LatencyToDatacenter int
	NumCaches           int
	CacheLatencies      map[int]int
}

type Cache struct {
	RemainingCapacity int
	Videos            map[int]bool
}

type RequestBurst struct {
	EndpointID int
	VideoID    int
	Num        int
}
