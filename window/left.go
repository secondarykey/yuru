package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type LeftMenu struct {
	drops *SelectDrops
}

type Operation int
type OperationMode int

const (
	MoveMode OperationMode = iota
	EditMode
)

const (
	LeftMenuMarginX     = 10
	LeftMenuMarginY     = 10
	LeftMenuDropMarginY = 10
)

func NewLeftMenu() (*LeftMenu, error) {

	var rtn LeftMenu

	x := float64(LeftMenuMarginX)
	y := float64(LeftMenuMarginY)

	rtn.drops = NewSelectDrops()

	for d := 0; d < int(None); d = d + 10 {
		d, err := NewDrop(d, int(x), int(y), DropWidth, DropHeight)
		if err != nil {
			return nil, xerrors.Errorf("NewDrop() error: %w", err)
		}
		rtn.drops.add(d)
		y = y + LeftMenuDropMarginY + DropHeight
	}

	return &rtn, nil
}

func (m *LeftMenu) Update(input *Input) error {
	if m.in(input) {
		return nil
	}
	return DoNotUpdate
}

func (m *LeftMenu) in(input *Input) bool {
	if LeftMenuMarginX < input.x && input.x < (DropWidth+LeftMenuMarginX) {
		if m.drops.in(input) {
			return true
		}
	}
	return false
}

func (m *LeftMenu) Draw(back *ebiten.Image) error {
	err := m.drops.Draw(back)
	if err != nil {
		return xerrors.Errorf("SelectDrops error: %w", err)
	}
	return nil
}

func (m LeftMenu) Get() Operation {
	return Operation(m.drops.getActive() * 10)
}

func (op Operation) GetMode() OperationMode {
	if op == 0 {
		return MoveMode
	}
	return EditMode
}
