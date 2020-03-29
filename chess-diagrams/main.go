package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Loading images ...")
	loadPieceImages()
	fmt.Println("Loading TrueType font ...")
	loadFont()

	fmt.Println("Starting server ...")
	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	log.Println("CustomerServer: Listening on http://localhost:8080/ ...")
	log.Fatal(http.ListenAndServe(":8080", router))

	/*
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	*/
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Request ...")

	img := drawDiagramForFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	png.Encode(w, img)
}
