package window

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/secondarykey/yuru/dto"
	"golang.org/x/xerrors"
)

type BoardMode int

const (
	DisplayBoardMode = iota
	EditBoardMode
)

type Board struct {
	data  dto.Board
	drops *BoardDrops

	startX int
	startY int

	mode BoardMode
}

func NewBoard(board dto.Board, startX, startY int, mode BoardMode) (*Board, error) {

	var b Board
	b.startX = startX
	b.startY = startY
	b.mode = mode
	b.data = board

	err := b.initBoard()
	if err != nil {
		return nil, xerrors.Errorf("initBoard() error: %w", err)
	}

	return &b, nil
}

func (b *Board) initBoard() error {
	b.drops = NewBoardDrops()
	y := b.startY
	for r := range b.data {
		row := b.data[r]
		x := b.startX
		for c := range row {
			d, err := NewDrop(row[c], int(x), int(y), DropWidth, DropHeight)
			if err != nil {
				return xerrors.Errorf("NewDrop() error: %w", err)
			}
			b.drops.add(d)
			x = x + DropWidth
		}
		y = y + DropHeight
	}
	return nil
}

func (b *Board) in(x, y int) bool {
	if b == nil {
		return false
	}
	rows := len(b.data)
	cols := len(b.data[0])

	h := rows*DropHeight + b.startY
	w := cols*DropWidth + b.startX

	if b.startX < x && x < w {
		if b.startY < y && y < h {
			return true
		}
	}
	return false
}

func (b *Board) isDragging() bool {
	if b == nil {
		return false
	}
	if b.drops.drags != nil {
		return true
	}
	if len(b.drops.moves) > 0 {
		return true
	}
	return false
}

func (b *Board) Update(input *Input, ope Operation) error {

	if b == nil {
		return nil
	}

	err := b.drops.Update(input, b.mode, ope)
	if err != nil {
		return xerrors.Errorf("drops Update() error: %w", err)
	}

	return nil
}

func (b *Board) Draw(back *ebiten.Image) error {

	if b == nil {
		return nil
	}

	err := b.drops.Draw(back)
	if err != nil {
		return xerrors.Errorf("BoardDrops.Draw() error: %w", err)
	}

	return nil
}

func (b *Board) export() (dto.Board, error) {

	var rtn dto.Board
	if b.drops.Lock() {
		return rtn, fmt.Errorf("drop lock error")
	}

	rtn = make([][]int, b.data.R())

	y := b.startY + DropHeight/2
	ydx := 0

	for ydx < b.data.R() {

		rtn[ydx] = make([]int, b.data.C())
		xdx := 0
		x := b.startX + DropWidth/2

		for xdx < b.data.C() {
			for _, drop := range b.drops.drops {
				if drop.Sprite.in(x, y) {
					rtn[ydx][xdx] = int(drop.Type)
					xdx++
					x = x + DropWidth
					break
				}
			}
		}

		y = y + DropHeight
		ydx++
	}

	return rtn, nil
}

func (b *Board) reset() error {
	err := b.initBoard()
	if err != nil {
		return xerrors.Errorf("initBoard() error: %w", err)
	}
	return nil
}
