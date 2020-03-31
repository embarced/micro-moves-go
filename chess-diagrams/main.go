package chessdiagrams

import (
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	startingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	emptyFen    = "8/8/8/8/8/8/8/8 w KQkq - 0 1"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/board.png", boardHandler)
	router.HandleFunc("/", defaultHandler)

	log.Println("chess-diagrams: Listening on http://localhost:8080/ ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	var fen string

	fens, ok := r.URL.Query()["fen"]
	if !ok || len(fens[0]) < 1 {
		fen = emptyFen
	} else {
		fen = fens[0]
	}

	img := drawDiagramForFen(fen)
	png.Encode(w, img)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	infile, _ := os.Open("html/index.html")
	io.Copy(w, infile)
}
