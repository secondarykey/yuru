package main

import (
	"container/heap"
	"sync"

	"github.com/cheggaaa/pb"
)

// T = Turn
// B = Beam
func search(T, B int) *State {

	num := count(G)

	maxCount := T * B * R * C
	bar := pb.StartNew(maxCount)

	//wg := &sync.WaitGroup{}

	//ch := make(chan *State)

	//開始点を設定
	for sr := 0; sr < R; sr++ {
		for sc := 0; sc < C; sc++ {

			//go func() {

			//wg.Add(1)

			q := make(Queue, 0)
			initial := NewState(num, sr, sc, 0, nil, G)

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

						if nr < 0 || R <= nr || nc < 0 || C <= nc {
							continue
						}

						if len(curRoute) != 0 && ((curRoute[len(curRoute)-1]+N/2)%N) == i {
							continue
						}

						nsRoute := append(curRoute, i)

						ns := NewState(0, nr, nc, turn+1, nsRoute, curG)
						ns.G[curR][curC], ns.G[nr][nc] = ns.G[nr][nc], ns.G[curR][curC]

						ns.combo = count(ns.G)
						heap.Push(&nq, ns)
					}
				}

				q = nq
				heap.Push(&bestQ, nq[0])
			}

			best := bestQ.Pop().(*State)
			best.startR = sr
			best.startC = sc

			//ch <- best
			//wg.Done()
			//}()

		}
	}

	bar.Finish()
	//wg.Wait()

	res := NewState(num, -1, -1, 0, nil, G)
	for {
		state, ok := <-ch
		if !ok {
			break
		}
		if !res.Less(state) {
			res = state
		}
	}
	return res
}

func count(p Board) int {

	comboMap := createComboMap(p)

	res := 0
	for _, v := range comboMap {

		seen := make([][]bool, R)
		for r := 0; r < R; r++ {
			seen[r] = make([]bool, C)
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

		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				if seen[r][c] {
					res++
					dfs(r, c, seen)
				}
			}
		}
	}

	return res
}

func createComboMap(p Board) map[int][]*Combo {

	rtnMap := make(map[int][]*Combo)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if c == 0 || p[r][c] != p[r][c-1] {
				nya(p, r, c, r, c, 0, 1, rtnMap)
			}
		}
	}

	for c := 0; c < C; c++ {
		for r := 0; r < R; r++ {
			if r == 0 || p[r][c] != p[r-1][c] {
				nya(p, r, c, r, c, 1, 0, rtnMap)
			}
		}
	}

	return rtnMap
}

func nya(p Board, sr, sc, cr, cc, dr, dc int, rtnMap map[int][]*Combo) {

	nr := cr + dr
	nc := cc + dc

	if nr < R && nc < C && p[nr][nc] == p[sr][sc] {
		nya(p, sr, sc, nr, nc, dr, dc, rtnMap)
	} else {
		dist := (cr-sr+1)*dr + (cc-sc+1)*dc

		elm := 3
		//2 Way
		if p[sr][sc] == 7 {
			elm = 4
		}

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

func dfs(r, c int, seen [][]bool) {

	seen[r][c] = false
	for i := 0; i < N; i++ {
		nr := r + DR[i]
		nc := c + DC[i]

		if nr < 0 || R <= nr || nc < 0 || C <= nc {
			continue
		}

		if seen[nr][nc] {
			dfs(nr, nc, seen)
		}
	}
}
