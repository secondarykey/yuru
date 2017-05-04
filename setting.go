package main

import ()

const (
	R         = 5
	C         = 6
	DIRECTION = "NESW"
)

var G Board = Board{
	{0, 2, 0, 3, 0, 5},
	{3, 0, 1, 2, 0, 0},
	{2, 1, 5, 0, 5, 5},
	{2, 5, 5, 2, 0, 5},
	{0, 4, 3, 3, 2, 4},
}

//DIRECTION
var DR [4]int = [4]int{-1, 0, 1, 0}
var DC [4]int = [4]int{0, 1, 0, -1}

var N int = len(DR)
