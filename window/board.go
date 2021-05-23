package window

import "github.com/hajimehoshi/ebiten/v2"

type Board struct {
}

func NewBoard() (*Board, error) {
	var b Board
	return &b, nil
}

func (b *Board) Update(input *Input) error {
	return nil
}

func (b *Board) Draw(back *ebiten.Image) error {

	return nil
}
