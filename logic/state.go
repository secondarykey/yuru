package logic

import (
	"bytes"
	"fmt"

	"github.com/secondarykey/yuru/dto"
)

type State struct {
	combo int
	nowR  int
	nowC  int
	route Route
	G     dto.Board
	turn  int

	startR int
	startC int
}

type Queue []*State

func NewState(r, c, turn int, route Route, P dto.Board) *State {

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
func (s *State) Max(max int) bool {
	return max == s.combo
}

func (s *State) Less(t *State) bool {
	if s.combo > t.combo {
		return true
	} else if s.combo == t.combo && s.turn < t.turn {
		return true
	}
	return false
}

func (s *State) PositionString() string {
	return fmt.Sprintf("Start(%d,%d)-End(%d,%d)", s.startR+1, s.startC+1, s.nowR+1, s.nowC+1)
}

func (s *State) Combo() int {
	return s.combo
}

func (s *State) RouteString() string {
	return s.route.DirectionString()
}

func (s *State) String() string {
	rtn := bytes.NewBuffer(make([]byte, 0, 200))

	rtn.WriteString(s.PositionString())
	rtn.WriteString(fmt.Sprintln())
	rtn.WriteString(s.route.String())
	rtn.WriteString(fmt.Sprintln())
	rtn.WriteString(fmt.Sprintf("combo:%d", s.combo))

	rtn.WriteString(fmt.Sprintln())
	rtn.WriteString(fmt.Sprintln())
	rtn.WriteString(s.G.String())
	return rtn.String()
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
