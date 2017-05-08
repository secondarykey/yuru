package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func init() {
	fmt.Println("初期盤面------------------")

}

func main() {

	cmds := os.Args
	var conf string
	if len(cmds) > 2 {
		conf = cmds[1]
	}

	//ファイルの指定
	if conf == "" {
		conf = "yuru.xml"
	}

	err := initialize(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	G.Print()
	calcMax()

	rtn := search(gConf.Turn, gConf.Beam)
	rtn.Print()

	//再検索を行うかを判定

	if !max(rtn.combo) {
		fmt.Println("最大コンボが見つからなかったので、再検索するとかも可能にする")
		//rtn = search(BT*2, BB*2)
		//rtn.Print()
	}
}

// T = Turn , B = Beam
func search(T, B int) *State {

	fmt.Printf("Turn:%d-Beam:%d-------------\n", T, B)
	fmt.Println(time.Now())
	res := NewState(-1, -1, 0, nil, G)

	wg := &sync.WaitGroup{}
	ch := make(chan *State, gConf.Board.R*gConf.Board.C)

	for sr := 0; sr < gConf.Board.R; sr++ {
		for sc := 0; sc < gConf.Board.C; sc++ {
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
