package logic

import (
	"container/heap"
	"sync"

	"github.com/secondarykey/yuru/config"
	"github.com/secondarykey/yuru/dto"
)

const (
	DIRECTION = "NESW"
	DONE      = -1
)

var (
	DR [4]int = [4]int{-1, 0, 1, 0}
	DC [4]int = [4]int{0, 1, 0, -1}
	N  int    = len(DR)
)

func Max() int {

	conf := config.Get()

	max := 0
	for d := 0; d < 10; d++ {
		n := 0
		for r := 0; r < conf.Board.R; r++ {
			for c := 0; c < conf.Board.C; c++ {
				if conf.BoardData[r][c] == d {
					n++
				}
			}
		}
		max += +n / 3
	}
	return max
}

// T = Turn , B = Beam
func Search() (*State, error) {

	conf := config.Get()

	res := NewState(-1, -1, 0, nil, conf.BoardData)

	T := conf.Turn
	B := conf.Beam

	wg := &sync.WaitGroup{}
	ch := make(chan *State, conf.Board.R*conf.Board.C)

	startR := 0
	startC := 0
	endR := conf.Board.R
	endC := conf.Board.C

	if conf.StartR > 0 {
		startR = conf.StartR - 1
		endR = startR + 1
	}
	if conf.StartC > 0 {
		startC = conf.StartC - 1
		endC = startC + 1
	}

	for sr := startR; sr < endR; sr++ {
		for sc := startC; sc < endC; sc++ {
			wg.Add(1)
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

	return res, nil
}

// T = Turn
// B = Beam
func analysis(T, B, r, c int, wg *sync.WaitGroup, ch chan *State) {

	conf := config.Get()

	q := make(Queue, 0)
	initial := NewState(r, c, 0, nil, conf.BoardData.Copy())

	heap.Push(&q, initial)
	bestQ := make(Queue, 0)

	for turn := 0; turn < T; turn++ {

		nq := make(Queue, 0)
		for k := 0; k < B; k++ {
			if len(q) == 0 {
				break
			}
			cur := q.Pop().(*State)

			curR := cur.nowR
			curC := cur.nowC
			curRoute := cur.route
			curG := cur.G

			for i := 0; i < N; i++ {

				nr := curR + DR[i]
				nc := curC + DC[i]

				if nr < 0 || conf.Board.R <= nr || nc < 0 || conf.Board.C <= nc {
					continue
				}

				if len(curRoute) != 0 && ((curRoute[len(curRoute)-1]+N/2)%N) == i {
					continue
				}

				nsRoute := append(curRoute.Copy(), i)
				tG := curG.Copy()
				tG[curR][curC], tG[nr][nc] = tG[nr][nc], tG[curR][curC]

				ns := NewState(nr, nc, turn+1, Route(nsRoute), tG)
				heap.Push(&nq, ns)
			}
		}

		q = nq
		heap.Push(&bestQ, nq[0])
	}

	best := bestQ.Pop().(*State)
	best.startR = r
	best.startC = c

	ch <- best
	wg.Done()
}

func count(p dto.Board) int {

	conf := config.Get()
	comboMap := createComboMap(p)

	d := p.Copy()
	res := 0
	for _, v := range comboMap {

		seen := make([][]bool, conf.Board.R)
		for r := 0; r < conf.Board.R; r++ {
			seen[r] = make([]bool, conf.Board.C)
		}

		for _, combo := range v {
			if combo.direction {
				for r := combo.startR; r <= combo.endR; r++ {
					seen[r][combo.startC] = true
				}
			} else {
				for c := combo.startC; c <= combo.endC; c++ {
					seen[combo.startR][c] = true
				}
			}
		}

		for r := 0; r < conf.Board.R; r++ {
			for c := 0; c < conf.Board.C; c++ {
				if seen[r][c] {
					res++
					dfs(r, c, seen, d)
				}
			}
		}
	}

	if res != 0 {
		res += down(d)
	}

	return res
}

func createComboMap(p dto.Board) map[int][]*Combo {

	conf := config.Get()
	rtnMap := make(map[int][]*Combo)

	for r := 0; r < conf.Board.R; r++ {
		for c := 0; c < conf.Board.C; c++ {
			if (c == 0 || p[r][c] != p[r][c-1]) && p[r][c] != DONE {
				nya(p, r, c, r, c, 0, 1, rtnMap)
			}
		}
	}

	for c := 0; c < conf.Board.C; c++ {
		for r := 0; r < conf.Board.R; r++ {
			if (r == 0 || p[r][c] != p[r-1][c]) && p[r][c] != DONE {
				nya(p, r, c, r, c, 1, 0, rtnMap)
			}
		}
	}

	return rtnMap
}

func nya(p dto.Board, sr, sc, cr, cc, dr, dc int, rtnMap map[int][]*Combo) {

	conf := config.Get()
	nr := cr + dr
	nc := cc + dc

	if nr < conf.Board.R && nc < conf.Board.C && p[nr][nc] == p[sr][sc] {
		nya(p, sr, sc, nr, nc, dr, dc, rtnMap)
	} else {
		dist := (cr-sr+1)*dr + (cc-sc+1)*dc

		elm := 3
		//2 Way
		//if p[sr][sc] == 7 {
		//elm = 4
		//}

		if dist >= elm {

			if rtnMap[p[sr][sc]] == nil {
				rtnMap[p[sr][sc]] = make([]*Combo, 0)
			}

			status := Combo{
				startR:    sr,
				startC:    sc,
				endR:      cr,
				endC:      cc,
				direction: dr == 1,
			}
			rtnMap[p[sr][sc]] = append(rtnMap[p[sr][sc]], &status)
		}
	}
}

func dfs(r, c int, seen [][]bool, p dto.Board) {

	conf := config.Get()
	seen[r][c] = false
	p[r][c] = DONE

	for i := 0; i < N; i++ {

		nr := r + DR[i]
		nc := c + DC[i]

		if nr < 0 || conf.Board.R <= nr || nc < 0 || conf.Board.C <= nc {
			continue
		}

		if seen[nr][nc] {
			dfs(nr, nc, seen, p)
		}
	}
}

func down(p dto.Board) int {
	conf := config.Get()
	for c := 0; c < conf.Board.C; c++ {
		for d := 0; d < conf.Board.R; d++ {
			for r := 0; r < conf.Board.R-1; r++ {
				if p[r+1][c] == DONE {
					p[r][c], p[r+1][c] = p[r+1][c], p[r][c]
				}
			}
		}
	}
	return count(p)
}
