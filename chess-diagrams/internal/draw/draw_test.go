package draw

import (
	_ "github.com/micro-moves-go/chess-diagrams/internal/testing"
	"testing"
)

const startingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func TestDrawDiagramForFen(t *testing.T) {

	fen := startingPosition
	img := DrawDiagramForFen(fen)

	size := img.Bounds().Size()
	if size.X != BoardSize {
		t.Error("wrong diagram size X")
	}
	if size.Y != BoardSize {
		t.Error("wrong diagram size Y")
	}
}

func BenchmarkDrawDiagramForFen(b *testing.B) {
	fen := startingPosition
	for i := 0; i < b.N; i++ {
		DrawDiagramForFen(fen)
	}
}
