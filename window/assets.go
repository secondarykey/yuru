package window

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/xerrors"
)

//go:embed _assets
var embAssets embed.FS
var assets fs.FS

var imageCache map[string]image.Image

func init() {
	var err error
	assets, err = fs.Sub(embAssets, "_assets")
	if err != nil {
		log.Println("assets Sub() error:", err)
	}

	imageCache = make(map[string]image.Image)

	err = initFont()
	if err != nil {
		log.Println("loadFont() error:", err)
	}
}

func convertImage(img image.Image) *ebiten.Image {
	return ebiten.NewImageFromImage(img)
}

func LoadDropImage(id int) (*ebiten.Image, error) {

	if id == int(Empty) {
		img := ebiten.NewImage(DropWidth, DropHeight)
		c := color.White
		w := float64(DropWidth)
		h := float64(DropHeight)

		//cross icon
		ebitenutil.DrawLine(img, 0, 0, w, 0, c)
		ebitenutil.DrawLine(img, w, 0, w, h, c)
		ebitenutil.DrawLine(img, w, h, 0, h, c)
		ebitenutil.DrawLine(img, 0, h, 0, 0, c)
		ebitenutil.DrawLine(img, 0, h/2, w, h/2, c)
		ebitenutil.DrawLine(img, w/2, 0, w/2, h, c)

		return img, nil
	}

	name := fmt.Sprintf("%03d.png", id)
	img, err := loadImage(name)
	if err != nil {
		return nil, xerrors.Errorf("LoadImage() error: %w", err)
	}
	return convertImage(img), nil
}

func loadImage(name string) (image.Image, error) {

	img, ok := imageCache[name]
	if ok {
		return img, nil
	}

	f, err := assets.Open(name)
	if err != nil {
		return nil, xerrors.Errorf("assets.Open() error: %w", err)
	}

	img, err = png.Decode(f)
	if err != nil {
		return nil, xerrors.Errorf("png.Decode() error: %w", err)
	}

	imageCache[name] = img

	return img, nil
}

const defaultFontName = "OdibeeSans-Regular.ttf"
const defaultFontDPI = 72
const defaultFontSize = 20
const smallFontSize = 20

var defaultFont font.Face
var smallFont font.Face

func initFont() error {

	f, err := assets.Open(defaultFontName)
	if err != nil {
		return xerrors.Errorf("assets.Open() error: %w", err)
	}

	info, err := f.Stat()
	if err != nil {
		return xerrors.Errorf("file.Stat() error: %w", err)
	}

	sz := info.Size()
	data := make([]byte, sz)
	_, err = f.Read(data)
	if err != nil {
		return xerrors.Errorf("file.Read() error: %w", err)
	}

	font, err := opentype.Parse(data)
	if err != nil {
		return xerrors.Errorf("file.Read() error: %w", err)
	}

	defaultFont, err = opentype.NewFace(font, &opentype.FaceOptions{
		Size: defaultFontSize,
		DPI:  defaultFontDPI,
	})

	smallFont, err = opentype.NewFace(font, &opentype.FaceOptions{
		Size: smallFontSize,
		DPI:  defaultFontDPI,
	})

	return nil
}
