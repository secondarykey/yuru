package window

import (
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	*Sprite
}

type TextBox struct {
	*Sprite
}

func NewTextBox(txt string, x, y int) (*TextBox, error) {

	var rtn TextBox

	//一番大きな文字列を幅に設定
	//何行あるかで高さを決定
	lines := strings.Split(txt, "\n")
	num := len(lines)

	w := 0
	for _, line := range lines {
		tw := font.MeasureString(smallFont, line).Ceil()
		if tw > w {
			w = tw
		}
	}

	fh := smallFont.Metrics().XHeight.Ceil()
	th := fh + 8
	h := th * num

	img := ebiten.NewImage(w, h)
	//img.Fill(color.White)

	dy := th

	for _, line := range lines {
		text.Draw(img, line, smallFont, 0, dy, color.RGBA{0, 255, 0, 255})
		dy += th
	}

	rtn.Sprite = NewSprite(x, y, w, h, img)
	return &rtn, nil
}
