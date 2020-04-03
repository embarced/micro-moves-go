package main

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
	FontPath = "fonts"
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

func drawKey(img *image.RGBA, col color.Color, squareSize, startX, startY int) {

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

	rankLabels := "87654321"
	pos_x_1 := startX/2 - int(fontSize/2.)
	pos_x_2 := pos_x_1 + int(8.5*float64(squareSize))
	posY := startY + int(fontSize+float64(squareSize)/6.)

	for _, c := range rankLabels {
		label := string(c)
		d.Dot = fixed.P(pos_x_1, posY)
		d.DrawString(label)
		d.Dot = fixed.P(pos_x_2, posY)
		d.DrawString(label)
		posY += squareSize
	}

	fileLabels := "abcdefgh"
	posX := startX + int(squareSize/4)
	pos_y_1 := startY - int(fontSize/1.5)
	pos_y_2 := pos_y_1 + int(8.5*float64(squareSize))
	for _, c := range fileLabels {
		label := string(c)
		d.Dot = fixed.P(posX, pos_y_1)
		d.DrawString(label)
		d.Dot = fixed.P(posX, pos_y_2)
		d.DrawString(label)
		posX += squareSize
	}
}
