package main

import (
	"fmt"
)

type Board [][]int

const (
	HEART = iota
	RED
	BLUE
	GREEN
	YELLOW
	BLACK
	DISTURB
	POISON
	DEADLY
	BOMB
)

const DONE = -1

func (b Board) Copy() Board {
	rtn := make(Board, len(b))

	for idx, elm := range b {
		rtn[idx] = make([]int, len(elm))
		copy(rtn[idx], elm)
	}
	return rtn
}

func (b Board) Print() {

	fmt.Println("    1 2 3 4 5 6")
	fmt.Println("------------------")
	for r := range b {
		fmt.Printf("%d | ", r+1)

		for c := range b[r] {
			fmt.Printf("%d ", b[r][c])
		}
		fmt.Println("")
	}
	fmt.Println("------------------")
}
