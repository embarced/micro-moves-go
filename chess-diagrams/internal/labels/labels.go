package labels

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"io/ioutil"
	"log"
)

const (
	FontPath = "assets/fonts"
	FontFile = "arial.ttf"
)

var (
	ttFont *truetype.Font
)

func init() {
	filename := FontPath + "/" + FontFile
	log.Printf("Loading font from [%s] ...\n", filename)

	fontBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	ttFont, err = truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
}

func DrawKey(img *image.RGBA, col color.Color, squareSize, startX, startY int) {

	fontSize := float64(squareSize) / 2.5

	d := &font.Drawer{
		Dst: img,
		Src: &image.Uniform{col},
		Face: truetype.NewFace(ttFont, &truetype.Options{
			Size:    fontSize,
			DPI:     72,
			Hinting: font.HintingNone,
		}),
	}

	const rankLabels = "87654321"
	posX1 := startX/2 - int(fontSize/2.)
	posX2 := posX1 + int(8.5*float64(squareSize))
	posY := startY + int(fontSize+float64(squareSize)/6.)

	for _, c := range rankLabels {
		label := string(c)
		d.Dot = fixed.P(posX1, posY)
		d.DrawString(label)
		d.Dot = fixed.P(posX2, posY)
		d.DrawString(label)
		posY += squareSize
	}

	const fileLabels = "abcdefgh"
	posX := startX + int(squareSize/4)
	posY1 := startY - int(fontSize/1.5)
	posY2 := posY1 + int(8.5*float64(squareSize))
	for _, c := range fileLabels {
		label := string(c)
		d.Dot = fixed.P(posX, posY1)
		d.DrawString(label)
		d.Dot = fixed.P(posX, posY2)
		d.DrawString(label)
		posX += squareSize
	}
}
