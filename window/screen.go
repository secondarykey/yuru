package window

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type Screen struct {
	image *ebiten.Image
	left  *LeftMenu
	board *Board
}

var screenColor = color.RGBA{0xbb, 0xad, 0xa0, 0xff}

func NewScreen(w, h int) (*Screen, error) {

	var s Screen
	var err error

	s.image = ebiten.NewImage(w, h)
	//s.image.Fill(screenColor)

	s.left, err = NewLeftMenu()
	if err != nil {
		return nil, xerrors.Errorf("NewLeftMenu() error: %w", err)
	}

	s.board, err = NewBoard()
	if err != nil {
		return nil, xerrors.Errorf("NewBoard() error: %w", err)
	}

	return &s, nil
}

var DoNotUpdate = fmt.Errorf("Do Not Update")

func (s *Screen) Update(input *Input) error {

	ebiten.SetCursorShape(ebiten.CursorShapeDefault)

	err := s.left.Update(input)
	if err != nil && !errors.Is(err, DoNotUpdate) {
		return xerrors.Errorf("LeftMenu error: %w", err)
	}

	err = s.board.Update(input)
	if err != nil && !errors.Is(err, DoNotUpdate) {
		return xerrors.Errorf("Board error: %w", err)
	}

	return DoNotUpdate
}

func (s *Screen) Draw(image *ebiten.Image) error {
	s.left.Draw(image)
	s.board.Draw(image)
	return nil
}

func (s *Screen) Get() *ebiten.Image {
	return s.image
}
