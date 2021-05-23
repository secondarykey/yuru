package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type DropType int

// ドロ強
// ロック
// 暗闇

const (
	Empty      DropType = 0
	Fire       DropType = 10
	Water      DropType = 20
	Wood       DropType = 30
	Light      DropType = 40
	Dark       DropType = 50
	Heart      DropType = 60
	Disturb    DropType = 70
	Poison     DropType = 80
	VeryPoison DropType = 90
	Bomb       DropType = 100
	None       DropType = 110
)

type Drop struct {
	Type DropType
	*Sprite
}

const (
	DropWidth  = 32
	DropHeight = 32
)

func NewDrop(t int, x, y, w, h int) (*Drop, error) {

	var d Drop

	img, err := LoadDropImage(t)
	if err != nil {
		return nil, xerrors.Errorf("LoadDropImage() error: %w", err)
	}

	d.Type = DropType(t)
	d.Sprite = NewSprite(x, y, w, h, img)

	return &d, nil
}

type SelectDrops struct {
	active int
	focus  int
	drops  []*Drop
}

func NewSelectDrops() *SelectDrops {
	var rtn SelectDrops
	rtn.active = 0
	rtn.focus = -1
	rtn.drops = make([]*Drop, 0)
	return &rtn
}

func (s *SelectDrops) add(d *Drop) {
	s.drops = append(s.drops, d)
}

func (s *SelectDrops) in(input *Input) bool {
	s.focus = -1
	for idx, d := range s.drops {
		sp := d.Sprite
		if sp.in(input.x, input.y) {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
			s.focus = idx
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				s.active = idx
			}
			return true
		}
	}
	return false
}

func (s *SelectDrops) Draw(back *ebiten.Image) error {

	for idx, d := range s.drops {
		sp := d.Sprite
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.x), float64(sp.y))

		if s.active == idx {
			op.ColorM.Scale(1, 1, 1, 0.3)
		} else if s.focus == idx {
			op.ColorM.Scale(1, 1, 1, 0.7)
		}
		back.DrawImage(sp.image, op)
	}
	return nil
}
