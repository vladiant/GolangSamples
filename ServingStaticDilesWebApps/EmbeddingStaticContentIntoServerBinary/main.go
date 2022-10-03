package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed public
var public embed.FS

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("02 Jan 2006 15:04:05 MST"))
}

func main() {
	publicFS, err := fs.Sub(public, "public")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/time", timeHandler)
	http.Handle("/", http.FileServer(http.FS(publicFS)))

	port := ":9999"
	log.Fatal(http.ListenAndServe(port, nil))
}
