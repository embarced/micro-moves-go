package main

import (
	"image/png"
	"log"
	"net/http"
)

func main() {
	log.Println("Loading images ...")
	loadPieceImages()
	log.Println("Loading TrueType font ...")
	loadFont()

	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	log.Println("CustomerServer: Listening on http://localhost:8080/ ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	img := drawDiagramForFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	png.Encode(w, img)
}
