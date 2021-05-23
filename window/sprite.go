package window

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	x     int
	y     int
	w     int
	h     int
	image *ebiten.Image
}

func NewSprite(x, y, w, h int, img *ebiten.Image) *Sprite {
	var s Sprite
	s.x = x
	s.y = y
	s.w = w
	s.h = h
	s.image = img
	return &s
}

func (s *Sprite) in(x, y int) bool {

	if s.x <= x && x <= (s.x+s.w) {
		if s.y <= y && y <= (s.y+s.h) {
			return true
		}
	}
	return false
}
