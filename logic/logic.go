package logic

import (
	"container/heap"
	"sync"

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

func Max(board dto.Board) int {

	max := 0
	for d := 0; d < 10; d++ {
		n := 0
		for r := 0; r < board.R(); r++ {
			for c := 0; c < board.C(); c++ {
				if board[r][c] == d {
					n++
				}
			}
		}
		max += +n / 3
	}
	return max
}

// T = Turn , B = Beam
func Search(board dto.Board, T, B, startR, startC int) (*State, error) {

	res := NewState(-1, -1, 0, nil, board)

	wg := &sync.WaitGroup{}
	ch := make(chan *State, board.R()*board.C())

	sR := 0
	sC := 0
	eR := board.R()
	eC := board.C()

	if startR > 0 {
		sR = startR - 1
		eR = startR + 1
	}
	if startC > 0 {
		sC = startC - 1
		eC = startC + 1
	}

	for sr := sR; sr < eR; sr++ {
		for sc := sC; sc < eC; sc++ {
			wg.Add(1)
			go analysis(board, T, B, sr, sc, wg, ch)
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
func analysis(board dto.Board, T, B, r, c int, wg *sync.WaitGroup, ch chan *State) {

	q := make(Queue, 0)
	initial := NewState(r, c, 0, nil, board.Copy())

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

				if nr < 0 || board.R() <= nr || nc < 0 || board.C() <= nc {
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

	comboMap := createComboMap(p)

	d := p.Copy()
	res := 0
	for _, v := range comboMap {

		seen := make([][]bool, p.R())
		for r := 0; r < p.R(); r++ {
			seen[r] = make([]bool, p.C())
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

		for r := 0; r < p.R(); r++ {
			for c := 0; c < p.C(); c++ {
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

	rtnMap := make(map[int][]*Combo)

	for r := 0; r < p.R(); r++ {
		for c := 0; c < p.C(); c++ {
			if (c == 0 || p[r][c] != p[r][c-1]) && p[r][c] != DONE {
				nya(p, r, c, r, c, 0, 1, rtnMap)
			}
		}
	}

	for c := 0; c < p.C(); c++ {
		for r := 0; r < p.R(); r++ {
			if (r == 0 || p[r][c] != p[r-1][c]) && p[r][c] != DONE {
				nya(p, r, c, r, c, 1, 0, rtnMap)
			}
		}
	}

	return rtnMap
}

func nya(p dto.Board, sr, sc, cr, cc, dr, dc int, rtnMap map[int][]*Combo) {

	nr := cr + dr
	nc := cc + dc

	if nr < p.R() && nc < p.C() && p[nr][nc] == p[sr][sc] {
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

	seen[r][c] = false
	p[r][c] = DONE

	for i := 0; i < N; i++ {

		nr := r + DR[i]
		nc := c + DC[i]

		if nr < 0 || p.R() <= nr || nc < 0 || p.C() <= nc {
			continue
		}

		if seen[nr][nc] {
			dfs(nr, nc, seen, p)
		}
	}
}

func down(p dto.Board) int {
	for c := 0; c < p.C(); c++ {
		for d := 0; d < p.R(); d++ {
			for r := 0; r < p.R()-1; r++ {
				if p[r+1][c] == DONE {
					p[r][c], p[r+1][c] = p[r+1][c], p[r][c]
				}
			}
		}
	}
	return count(p)
}
