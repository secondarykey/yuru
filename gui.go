package yuru

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/secondarykey/yuru/window"
	"golang.org/x/xerrors"
)

func Show(name string) error {

	win, err := window.New()
	if err != nil {
		return xerrors.Errorf("window.New() error: %w", err)
	}

	err = ebiten.RunGame(win)
	if err != nil {
		return xerrors.Errorf("ebiten.RunGame() error: %w", err)
	}

	return nil
}
