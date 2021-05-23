package window

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	state PointerState

	x int
	y int
}

type PointerState int

const (
	PointerStateNeutral PointerState = iota
	PointerStateClick
	PointerStateDrag
	PointerStateDrop
)

func NewInput() *Input {
	var rtn Input
	rtn.state = PointerStateNeutral
	return &rtn
}

func (i *Input) Update() {
	i.x, i.y = ebiten.CursorPosition()

	switch i.state {
	case PointerStateNeutral:
	}

}
