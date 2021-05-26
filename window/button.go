package window

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	focus bool
	*Sprite
}

func NewButton(v string, x, y, w, h int) (*Button, error) {

	var btn Button
	img := ebiten.NewImage(w, h)
	img.Fill(color.RGBA{220, 220, 220, 255})

	tw := font.MeasureString(defaultFont, v).Ceil()
	th := defaultFont.Metrics().XHeight.Ceil()
	if tw > w {
		return nil, fmt.Errorf("font width over error: %d > %d", tw, w)
	} else if th > h {
		return nil, fmt.Errorf("font height over error: %d > %d", th, h)
	}

	centerX := w / 2
	centerY := h / 2

	dx := centerX - tw/2
	dy := centerY + th/2

	text.Draw(img, v, defaultFont, dx, dy, color.Black)

	btn.Sprite = NewSprite(x, y, w, h, img)
	btn.focus = false
	return &btn, nil
}
