package main

// Problem defines the input parameters
type Problem struct {
	NumRows    int
	NumColumns int
	L          int
	H          int
	Grid       [][]string
}

type Result []Slice

type Slice struct {
	row1, row2, column1, column2 int
}
