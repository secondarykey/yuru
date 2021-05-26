package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	mx   int
	my   int
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
	d.mx = x
	d.my = y

	return &d, nil
}

func (d *Drop) ChangeType(t int) error {
	img, err := LoadDropImage(t)
	if err != nil {
		return xerrors.Errorf("LoadDropImage() error: %w", err)
	}
	d.Type = DropType(t)
	d.Sprite = NewSprite(d.Sprite.x, d.Sprite.y, d.Sprite.w, d.Sprite.h, img)
	return nil
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

func (s *SelectDrops) getActive() int {
	return s.active
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

type BoardDrops struct {
	focus   int
	drags   *Drop
	moves   []*Drop
	unmoves []*Drop
	drops   []*Drop

	play *Play
}

func NewBoardDrops() *BoardDrops {
	var rtn BoardDrops
	rtn.drops = make([]*Drop, 0)
	return &rtn
}

func (b *BoardDrops) add(d *Drop) {
	b.drops = append(b.drops, d)
}

type DragState int

const (
	DragNone DragState = iota
	DragStart
	DragEnd
)

func (b *BoardDrops) Update(input *Input, mode BoardMode, ope Operation) error {

	var dragState DragState = DragNone
	var edit int
	edit = -1
	b.focus = -1

	if mode == EditBoardMode {

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if ope.GetMode() == MoveMode {
				dragState = DragStart
			} else if ope.GetMode() == EditMode {
				edit = int(ope)
			}
		} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && b.drags != nil {
			dragState = DragEnd
		} else {
			if b.drags == nil {
				dragState = DragNone
			}
		}

		for idx, d := range b.drops {
			sp := d.Sprite
			if sp.in(input.x, input.y) {

				ebiten.SetCursorShape(ebiten.CursorShapePointer)
				if b.drags == nil {
					b.focus = idx

					if edit > 0 {
						d.ChangeType(edit)
					}
				} else {
					if b.drags != d {

						b.play.move(d)

						d.mx = sp.x
						d.my = sp.y
						sp.x, b.drags.Sprite.x = b.drags.Sprite.x, sp.x
						sp.y, b.drags.Sprite.y = b.drags.Sprite.y, sp.y
						b.moves = append(b.moves, d)
					}
				}
				break
			}
		}

		//drag start
		if b.focus != -1 && dragState == DragStart {
			b.drags = b.drops[b.focus]
			if b.play == nil {
				b.play = NewPlay(b.drags)
			}
		}

		//drag position
		if b.drags != nil {
			b.drags.mx = input.x - (DropWidth / 2)
			b.drags.my = input.y - (DropHeight / 2)
		}

		if dragState == DragEnd {
			b.play.stopPlay()

			sp := b.drags.Sprite
			//自分のあった位置に置く
			b.drags.mx, b.drags.my = sp.x, sp.y
			b.drags = nil
		}
	}

	newmove := make([]*Drop, 0)
	mp := 8
	for _, d := range b.moves {
		sp := d.Sprite
		mx := sp.x - d.mx
		if mx > 0 {
			d.mx = d.mx + mp
		} else if mx < 0 {
			d.mx = d.mx - mp
		}
		my := sp.y - d.my
		if my > 0 {
			d.my = d.my + mp
		} else if my < 0 {
			d.my = d.my - mp
		}
		if mx != 0 || my != 0 {
			newmove = append(newmove, d)
		}
	}
	b.moves = newmove

	b.unmoves = make([]*Drop, 0)
	for _, d := range b.drops {

		if b.drags == d {
			continue
		}
		un := true
		for _, m := range b.moves {
			if d == m {
				un = false
				break
			}
		}
		if un {
			b.unmoves = append(b.unmoves, d)
		}
	}

	return nil
}

func (b *BoardDrops) Draw(back *ebiten.Image) error {

	for idx, d := range b.unmoves {
		sp := d.Sprite
		op := &ebiten.DrawImageOptions{}
		x := float64(sp.x)
		y := float64(sp.y)
		if b.focus == idx {
			op.ColorM.Scale(1, 1, 1, 0.7)
		}

		op.GeoM.Translate(x, y)
		back.DrawImage(sp.image, op)
	}

	for _, d := range b.moves {
		op := &ebiten.DrawImageOptions{}
		x := float64(d.mx)
		y := float64(d.my)
		op.GeoM.Translate(x, y)
		back.DrawImage(d.Sprite.image, op)
	}

	if b.drags != nil {
		op := &ebiten.DrawImageOptions{}
		x := float64(b.drags.mx)
		y := float64(b.drags.my)
		op.GeoM.Translate(x, y)
		op.ColorM.Scale(1, 1, 1, 0.7)
		back.DrawImage(b.drags.Sprite.image, op)
	}

	return nil
}

func (b *BoardDrops) Lock() bool {
	if len(b.moves) > 0 {
		return true
	}
	return false
}

func (b *BoardDrops) GetPlay() *Play {
	return b.play
}
