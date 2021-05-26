package window

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/xerrors"
)

const (
	Width  = 600
	Height = 600
)

type Window struct {
	screen *Screen
	input  *Input
}

func New() (*Window, error) {
	var w Window
	var err error
	w.input = NewInput()
	w.screen, err = NewScreen(Width, Height)
	if err != nil {
		return nil, xerrors.Errorf("NewScreen() error: %w", err)
	}
	return &w, nil
}

func (win *Window) Layout(w, h int) (int, int) {
	return w, h
}

func (w *Window) Update() error {
	w.input.Update()
	err := w.screen.Update(w.input)
	if err != nil {
		if !errors.Is(err, DoNotUpdate) {
			return xerrors.Errorf("screen update error: %w", err)
		}
	}
	return nil
}

func (w *Window) Draw(s *ebiten.Image) {

	msg := fmt.Sprintf("TPS = %0.2f\nFPS = %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(s, msg)

	op := &ebiten.DrawImageOptions{}
	err := w.screen.Draw(s)
	if err != nil {
		log.Println(err)
	} else {
		s.DrawImage(w.screen.Get(), op)
	}

}
