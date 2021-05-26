package window

import (
	"github.com/hajimehoshi/ebiten/v2"
)

//Focus
//Active
type Component interface {
	Draw(*ebiten.Image) error
}

type Focuser interface {
	Focus()
}

type Pusher interface {
	Push()
}
