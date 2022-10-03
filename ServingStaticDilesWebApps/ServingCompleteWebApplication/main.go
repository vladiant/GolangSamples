package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format("02 Jan 2006 15:04:05 MST\n"))
}

func main() {
	port := ":9999"

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/time", timeHandler)

	http.ListenAndServe(port, nil)
}
