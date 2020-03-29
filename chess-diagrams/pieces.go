package main

import (
	"image"
	"image/png"
	"os"
	"unicode"
)

const (
	ImagePath = "images/pieces/32"
)

var (
	allPieceNames = []string{"bb", "bk", "bn", "bp", "bq", "br", "wb", "wk", "wn", "wp", "wq", "wr"}
	pieceImages   map[string]image.Image
)

func loadPieceImages() {
	pieceImages = make(map[string]image.Image)
	for _, name := range allPieceNames {
		img := loadPieceImage(name)
		pieceImages[name] = img
	}
}

func fenLetterToPieceName(char rune) string {
	if unicode.IsUpper(char) {
		return "w" + string(unicode.ToLower(char))
	} else {
		return "b" + string(char)
	}
}

func loadPieceImage(piece string) image.Image {
	infile, err := os.Open(ImagePath + "/" + piece + ".png")
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	return img
}
