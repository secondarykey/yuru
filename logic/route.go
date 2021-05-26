package logic

import (
	"bytes"
	"fmt"
	"strings"
)

type Route []int

func (r Route) Copy() []int {
	rtn := make([]int, len(r))
	copy(rtn, r)
	return rtn
}

func (r Route) DirectionString() string {
	var builder strings.Builder
	for _, elm := range r {
		w := "-"
		if elm >= 0 {
			w = string(DIRECTION[elm])
		}
		builder.WriteString(w)
	}
	return builder.String()
}

func (r Route) String() string {
	rtn := bytes.NewBuffer(make([]byte, 0, 100))
	rtn.WriteString(fmt.Sprintf("route:%d[", len(r)))
	rtn.WriteString(r.DirectionString())
	rtn.WriteString("]")
	return rtn.String()
}
