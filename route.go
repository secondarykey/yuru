package yuru

import (
	"fmt"
	"bytes"
)

type Route []int

func (r Route) Copy() []int {
	rtn := make([]int, len(r))
	copy(rtn, r)
	return rtn
}

func (r Route) String() string {
	rtn := bytes.NewBuffer(make([]byte,0,100))
	rtn.WriteString(fmt.Sprintf("route:%d[", len(r)))
	for _, elm := range r {
		rtn.WriteString(fmt.Sprintf(string(DIRECTION[elm])))
	}
	rtn.WriteString(fmt.Sprintf("]"))
	return rtn.String()
}
