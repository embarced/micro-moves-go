package main

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"io/ioutil"
)

const (
	FontPath = "fonts"
	FontFile = "arial.ttf"
)

var (
	ttFont *truetype.Font
)

func loadFont() {
	filename := FontPath + "/" + FontFile
	fontBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	ttFont, err = truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
}

func draw_key(img *image.RGBA, square_size, start_x, start_y int) {

	fontSize := float64(square_size) / 2.5
	font_dy := int((start_y - int(fontSize)) / 2)
	font_dx := int((start_x - int(fontSize/2)) / 2)

	d := &font.Drawer{
		Dst: img,
		Src: image.White,
		Face: truetype.NewFace(ttFont, &truetype.Options{
			Size:    fontSize,
			DPI:     72,
			Hinting: font.HintingNone,
		}),
	}
	d.Dot = fixed.P(100, 100)

	rankLabels := "87654321"
	pos_x_1 := font_dx
	pos_x_2 := font_dx + int(8.5*float64(square_size))
	pos_y := start_y + int(square_size/6)

	for _, c := range rankLabels {
		label := string(c)
		d.Dot = fixed.P(pos_x_1, pos_y)
		d.DrawString(label)
		d.Dot = fixed.P(pos_x_2, pos_y)
		d.DrawString(label)
		pos_y += square_size
	}

	fileLabels := "abcdefgh"
	pos_x := start_x + int(square_size/4) + font_dx
	pos_y_1 := start_y + font_dy
	pos_y_2 := start_y + 8*square_size + font_dy
	for _, c := range fileLabels {
		label := string(c)
		d.Dot = fixed.P(pos_x, pos_y_1)
		d.DrawString(label)
		d.Dot = fixed.P(pos_x, pos_y_2)
		d.DrawString(label)
		pos_x += square_size
	}
}
