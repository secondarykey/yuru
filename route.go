package main

import (
	"fmt"
)

type Route []int

func (r Route) Copy() []int {
	rtn := make([]int, len(r))
	copy(rtn, r)
	return rtn
}

func (r Route) Print() {
	fmt.Printf("route:%d[", len(r))
	for _, elm := range r {
		fmt.Printf(string(DIRECTION[elm]))
	}
	fmt.Printf("]")
}
