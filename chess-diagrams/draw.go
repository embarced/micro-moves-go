package main

import (
	"image"
	"image/color"
	"image/draw"
	"strconv"
	"strings"
	"unicode"
)

const (
	SquareSize = 32
	BorderSize = SquareSize / 2
	BoardSize  = SquareSize*8 + 2*BorderSize
)

var (
	BoardColorLight  = color.White
	BoardColorDark   = color.RGBA{211, 211, 211, 255} // lightgray
	BoardColorBorder = color.RGBA{66, 66, 66, 255}    // darkgray
	BoardColorKey    = color.White
)

func createImage(width, height int, c color.Color) image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	drawRectangle(*img, 0, 0, width, height, c)

	return *img
}

func drawRectangle(img image.RGBA, startX, startY, width, height int, c color.Color) {
	rect := image.Rectangle{
		Min: image.Point{startX, startY},
		Max: image.Point{startX + width, startY + height},
	}
	draw.Draw(&img, rect, &image.Uniform{c}, image.Point{}, draw.Src)
}

// Draw a checkered 8x8 chess board into a given images.
func drawBoard(img image.RGBA, squareSize int, startX int, startY int, light color.Color, dark color.Color) {
	var x, y int
	var fill color.Color
	for square := 0; square < 64; square++ {
		x, y = square%8, square/8
		if (x+y)%2 == 0 {
			fill = light
		} else {
			fill = dark
		}
		drawRectangle(img, startX+x*squareSize, startY+y*squareSize, squareSize, squareSize, fill)
	}
}

// Draws chess pieces into a given image.
func drawPieces(img image.RGBA, square_size, start_x, start_y int, pieces string) {

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

func drawPiece(img image.RGBA, square_size, posX, posY int, piece string) {
	pImg := pieceImages[piece]

	rect := image.Rect(posX, posY, posX+square_size, posY+square_size)
	point := image.Point{0, 0}
	draw.Draw(&img, rect, pImg, point, 0)
}

// Creates an image for the given FEN position.
//
func drawDiagramForFen(fen string) image.Image {

	img := createImage(BoardSize, BoardSize, BoardColorBorder)
	drawBoard(img, SquareSize, BorderSize, BorderSize, BoardColorLight, BoardColorDark)

	fenGroups := strings.Split(fen, " ")
	pieces := fenGroups[0]

	drawPieces(img, SquareSize, BorderSize, BorderSize, pieces)
	drawKey(&img, BoardColorKey, SquareSize, 20, 20)

	return &img
}
