package main

import (
	"fmt"
	"time"
)

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

var DR [4]int = [4]int{-1, 0, 1, 0}
var DC [4]int = [4]int{0, 1, 0, -1}

var N int = len(DR)

func init() {
	fmt.Println("初期盤面------------------")
	G.Print()
}

func main() {
	fmt.Println(time.Now())

	rtn := search(100, 500)
	fmt.Printf("簡易探査-Turn:%d-Beam:%d-------------\n", 100, 500)
	rtn.Print()

	fmt.Println(time.Now())

	turn := len(rtn.route)
	combo := rtn.combo

	rtn = search(turn, 5000)
	fmt.Printf("探査-Turn:%d-Beam:%d-------------\n", turn, 5000)
	rtn.Print()

	n := len(rtn.route)
	//良くなった部分があれば
	if n > combo || rtn.turn < turn {
		rtn = search(n, 20000)
		fmt.Printf("探査-Turn:%d-Beam:%d-------------\n", n, 20000)
		rtn.Print()
	}

	fmt.Println(time.Now())
}
