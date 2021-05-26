package window

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/secondarykey/yuru/config"
	"github.com/secondarykey/yuru/logic"
	"golang.org/x/xerrors"
)

type Screen struct {
	image *ebiten.Image
	left  *LeftMenu

	editBoard *Board
	play      *TextBox
	anaBtn    *Button
	resetBtn  *Button

	result       *TextBox
	displayBoard *Board
}

const (
	EditBoardStartX    = 70
	EditBoardStartY    = 20
	DisplayBoardStartX = 70
	DisplayBoardStartY = 250

	ResultStartX = 280
	ResultStartY = 250

	ButtonWidth     = 80
	ButtonHeight    = 32
	AnalysisButtonX = 125
	AnalysisButtonY = 200
	ResetButtonX    = 280
	ResetButtonY    = 20
)

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

	s.editBoard, err = NewBoard(config.GetDefaultBoard(), EditBoardStartX, EditBoardStartY, EditBoardMode)
	if err != nil {
		return nil, xerrors.Errorf("NewBoard() error: %w", err)
	}

	s.anaBtn, err = NewButton("Analysis", AnalysisButtonX, AnalysisButtonY, ButtonWidth, ButtonHeight)
	if err != nil {
		return nil, xerrors.Errorf("NewButton(analysis) error: %w", err)
	}

	s.resetBtn, err = NewButton("Reset", ResetButtonX, ResetButtonY, ButtonWidth, ButtonHeight)
	if err != nil {
		return nil, xerrors.Errorf("NewButton(reset) error: %w", err)
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

	err = s.editBoard.Update(input, s.left.Get())
	if err != nil && !errors.Is(err, DoNotUpdate) {
		return xerrors.Errorf("Board error: %w", err)
	}

	p := s.editBoard.drops.GetPlay()
	if p != nil {
		r := p.routeString()
		sr := splitRoute(r, 20, "\n", "    ")
		txt := fmt.Sprintf("Time:%ss\nRoute[%d]:\n%s", p.duration(), len(r), sr)

		s.play, err = NewTextBox(txt, 280, 100)
		if err != nil {
			return xerrors.Errorf("NewTextBox(play) error: %w", err)
		}
	}

	err = s.displayBoard.Update(input, s.left.Get())
	if err != nil && !errors.Is(err, DoNotUpdate) {
		return xerrors.Errorf("Board error: %w", err)
	}

	if s.anaBtn.in(input.x, input.y) {
		s.anaBtn.focus = true
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

			go func() {
				now, err := s.editBoard.export()
				if err != nil {
					log.Println(err)
					//return xerrors.Errorf("board export() error: %w", err)
					return
				}

				conf := config.Get()
				result, err := logic.Search(now, conf.Beam, conf.Turn, conf.StartR, conf.StartC)
				if err != nil {
					log.Println(err)
					//return xerrors.Errorf("logic.Search() error: %w", err)
					return
				}

				s.displayBoard, err = NewBoard(result.G, DisplayBoardStartX, DisplayBoardStartY, DisplayBoardMode)
				if err != nil {
					log.Println(err)
					//return xerrors.Errorf("NewBoard() error: %w", err)
					return
				}

				r := result.RouteString()
				route := splitRoute(r, 20, "\n", "    ")

				txt := result.PositionString() + "\n" +
					fmt.Sprintf("Combo:%d", result.Combo()) + "\n" +
					fmt.Sprintf("Route[%d]:\n%s", len(r), route)

				s.result, err = NewTextBox(txt, ResultStartX, ResultStartY)
				if err != nil {
					log.Println(err)
					//return xerrors.Errorf("NewTextBox(result) error: %w", err)
					return
				}
			}()

		}
	} else {
		s.anaBtn.focus = false
	}

	if s.resetBtn.in(input.x, input.y) {
		s.resetBtn.focus = true
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			s.result = nil
			s.displayBoard = nil
			err := s.editBoard.reset()
			if err != nil {
				return xerrors.Errorf("board reset() error: %w", err)
			}
		}
	} else {
		s.resetBtn.focus = false
	}

	return DoNotUpdate
}

func (s *Screen) Draw(image *ebiten.Image) error {

	s.left.Draw(image)
	s.editBoard.Draw(image)
	s.displayBoard.Draw(image)

	sp := s.anaBtn.Sprite
	op := &ebiten.DrawImageOptions{}
	if s.anaBtn.focus {
		op.ColorM.Scale(1, 1, 1, 0.7)
	}
	op.GeoM.Translate(float64(sp.x), float64(sp.y))
	image.DrawImage(sp.image, op)

	sp = s.resetBtn.Sprite
	op = &ebiten.DrawImageOptions{}
	if s.resetBtn.focus {
		op.ColorM.Scale(1, 1, 1, 0.7)
	}
	op.GeoM.Translate(float64(sp.x), float64(sp.y))
	image.DrawImage(sp.image, op)

	if s.result != nil {
		sp = s.result.Sprite
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.x), float64(sp.y))
		image.DrawImage(sp.image, op)
	}

	p := s.editBoard.drops.play
	if p.isStarted() || p.isStoped() {
		sp = s.play.Sprite
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.x), float64(sp.y))
		image.DrawImage(sp.image, op)
	}

	return nil
}

func (s *Screen) Get() *ebiten.Image {
	return s.image
}

func splitRoute(w string, l int, sep, indent string) string {

	var b strings.Builder
	c := len(w)
	start := 0
	end := l
	for {
		if end > c {
			end = c
		}

		b.WriteString(indent + w[start:end])

		if end == c {
			break
		}
		b.WriteString(sep)
		start, end = end, end+l
	}
	return b.String()
}
