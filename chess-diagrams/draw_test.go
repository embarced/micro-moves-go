package chessdiagrams

import (
	"testing"
)

func TestDrawDiagramForFen(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	img := drawDiagramForFen(fen)

	if img == nil {
		t.Fail()
	}
}
