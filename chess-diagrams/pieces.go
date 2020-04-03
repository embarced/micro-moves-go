package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	ImagePath = "images/pieces/32"
)

var (
	allPieceNames = []string{"bb", "bk", "bn", "bp", "bq", "br", "wb", "wk", "wn", "wp", "wq", "wr"}
	pieceImages   map[string]image.Image
)

func init() {
	log.Printf("Loading piece images from [%s] ...\n", ImagePath)

	pieceImages = make(map[string]image.Image)
	for _, name := range allPieceNames {
		img := loadPieceImage(name)
		pieceImages[name] = img
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

// Draws chess pieces into a given image.
func drawPieces(img *image.RGBA, square_size, start_x, start_y int, pieces string) {

	ranks := strings.Split(pieces, "/")

	for rank_no, rank := range ranks {
		file_no := 0
		for _, char := range rank {
			if unicode.IsDigit(char) {
				n, _ := strconv.Atoi(string(char))
				file_no += n
			} else {
				pieceName := fenLetterToPieceName(char)
				posX := start_x + square_size*file_no
				posY := start_y + square_size*rank_no
				drawPiece(img, square_size, posX, posY, pieceName)
				file_no += 1
			}
		}
	}

}

func drawPiece(img *image.RGBA, square_size, posX, posY int, piece string) {
	pImg := pieceImages[piece]

	rect := image.Rect(posX, posY, posX+square_size, posY+square_size)
	point := image.Point{0, 0}
	draw.Draw(img, rect, pImg, point, 0)
}

func fenLetterToPieceName(char rune) string {
	if unicode.IsUpper(char) {
		return "w" + string(unicode.ToLower(char))
	} else {
		return "b" + string(char)
	}
}
