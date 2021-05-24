package dto

import (
	"bytes"
	"fmt"
)

type Board [][]int

type BoardInfo struct {
	R int    `xml:"r,attr"`
	C int    `xml:"c,attr"`
	B string `xml:",chardata"`
}

const (
	HEART = iota * 10
	RED
	BLUE
	GREEN
	YELLOW
	BLACK
	DISTURB
	POISON
	DEADLY
	BOMB
)

//コピーで行う
func (b Board) Copy() Board {
	rtn := make(Board, len(b))

	for idx, elm := range b {
		rtn[idx] = make([]int, len(elm))
		copy(rtn[idx], elm)
	}
	return rtn
}

//盤面表示
func (b Board) String() string {

	rtn := bytes.NewBuffer(make([]byte, 0, 200))
	rtn.WriteString(fmt.Sprintln("     1  2  3  4  5  6"))
	rtn.WriteString(fmt.Sprintln("-----------------------"))
	for r := range b {
		rtn.WriteString(fmt.Sprintf("%d | ", r+1))

		for c := range b[r] {
			rtn.WriteString(fmt.Sprintf("%d ", b[r][c]))
		}
		rtn.WriteString(fmt.Sprintln(""))
	}
	rtn.WriteString(fmt.Sprintln("-----------------------"))
	return rtn.String()
}
