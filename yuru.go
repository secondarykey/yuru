package yuru

import (
	"sync"
)

// T = Turn , B = Beam
func Search(conf *Config) (*State,error) {

	res := NewState(-1, -1, 0, nil, conf.BoardData,conf)

	T := conf.Turn
	B := conf.Beam

	wg := &sync.WaitGroup{}
	ch := make(chan *State, conf.Board.R*conf.Board.C)

        startR := 0
        startC := 0
        endR := conf.Board.R
        endC := conf.Board.C

	if conf.StartR > 0 {
            startR = conf.StartR-1
            endR = startR + 1
	}
	if conf.StartC > 0 {
            startC = conf.StartC-1
            endC = startC + 1
        }

	for sr := startR; sr < endR; sr++ {
		for sc := startC; sc < endC; sc++ {
			wg.Add(1)
			go analysis(T, B, sr, sc, conf, wg, ch)
		}
	}

	wg.Wait()
	close(ch)

	for s := range ch {
		if !res.Less(s) {
			res = s
		}
	}

	return res,nil
}
