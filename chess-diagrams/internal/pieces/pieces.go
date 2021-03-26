package pieces

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
	ImagePath = "assets/images/pieces/32"
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
func DrawPieces(img *image.RGBA, squareSize, startX, startY int, pieces string) {

	ranks := strings.Split(pieces, "/")

	for rankNo, rank := range ranks {
		fileNo := 0
		for _, char := range rank {
			if unicode.IsDigit(char) {
				n, _ := strconv.Atoi(string(char))
				fileNo += n
			} else {
				pieceName := fenLetterToPieceName(char)
				posX := startX + squareSize*fileNo
				posY := startY + squareSize*rankNo
				drawPiece(img, squareSize, posX, posY, pieceName)
				fileNo++
			}
		}
	}

}

func drawPiece(img *image.RGBA, squareSize, posX, posY int, piece string) {
	pImg := pieceImages[piece]

	rect := image.Rect(posX, posY, posX+squareSize, posY+squareSize)
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
