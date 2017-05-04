package main

import (
	"fmt"
	"sync"
	"time"
)

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

// T = Turn , B = Beam
func search(T, B int) *State {

	res := NewState(-1, -1, 0, nil, G)

	wg := &sync.WaitGroup{}
	ch := make(chan *State, R*C)

	for sr := 0; sr < R; sr++ {
		for sc := 0; sc < C; sc++ {
			go analysis(T, B, sr, sc, wg, ch)
		}

	}

	wg.Wait()
	close(ch)

	for s := range ch {
		if !res.Less(s) {
			res = s
		}
	}
	return res
}
