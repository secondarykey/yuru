package main

import ()

const (
	R         = 5
	C         = 6
	DIRECTION = "NESW"
)

var G Board = Board{
	{2, 5, 5, 3, 2, 1},
	{4, 0, 5, 1, 5, 2},
	{1, 1, 2, 5, 5, 0},
	{3, 1, 4, 5, 4, 0},
	{5, 2, 0, 1, 0, 1},
}

var DR [4]int = [4]int{-1, 0, 1, 0}
var DC [4]int = [4]int{0, 1, 0, -1}

var N int = len(DR)
