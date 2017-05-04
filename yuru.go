package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("初期盤面------------------")

	G.Print()
	calcMax()
}

func main() {
	rtn := search(50, 50)
	rtn.Print()

	if !max(rtn.combo) {
		fmt.Println("最大コンボが見つかりませんでした")
		rtn = search(100, 100)
		rtn.Print()
	}
}

// T = Turn , B = Beam
func search(T, B int) *State {

	fmt.Printf("Turn:%d-Beam:%d-------------\n", T, B)
	fmt.Println(time.Now())
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

	fmt.Println(time.Now())
	return res
}
