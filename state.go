package main

import (
	"fmt"
)

type State struct {
	combo int
	nowR  int
	nowC  int
	route Route
	G     Board
	turn  int

	startR int
	startC int
}

type Queue []*State

func NewState(r, c, turn int, route Route, P Board) *State {

	s := State{
		nowR: r,
		nowC: c,
		turn: turn,
	}
	s.combo = count(P)

	if route == nil {
		s.route = make(Route, 0)
	} else {
		s.route = route
	}
	s.G = P

	return &s
}

func (s *State) Less(t *State) bool {
	if s.combo > t.combo {
		return true
	} else if s.combo == t.combo && s.turn < t.turn {
		return true
	}
	return false
}

func (s *State) Print() {

	s.G.Print()
	fmt.Printf("Start(%d,%d)-End(%d,%d)\n", s.startR+1, s.startC+1, s.nowR+1, s.nowC+1)
	s.route.Print()
	fmt.Println()
	fmt.Printf("combo:%d\n", s.combo)
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q Queue) Less(i, j int) bool {
	return q[i].Less(q[j])
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[0]
	*q = old[1:n]
	return item
}

func (q *Queue) Push(x interface{}) {
	item := x.(*State)
	*q = append(*q, item)
}
