package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var page = `
<html>
 <head>
   <link rel="stylesheet" type="text/css" href="/css/style.css">
 </head>
 <body>
   <h2>System clock (updates every second)</h2>
   <div id="output"></div>
   <script>
	let outputBox = document.querySelector('#output');

	window.addEventListener('DOMContentLoaded', (event) => {
		outputBox.innerHTML = "initializing...";
		tick();
		setInterval(tick, 1000);
	});
	
	function tick() {
		fetch('/time')
		.then((response) => {
			if (!response.ok) {
			throw new Error("error response");
			}
			return response.text();
		})
		.then((text) => {
			outputBox.innerHTML = text;
		})
		.catch((error) => {
			outputBox.innerHTML = "network error";
		});
	}
   </script>
 </body>
</html>
`

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if path := strings.Trim(r.URL.Path, "/"); len(path) > 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, page)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format("02 Jan 2006 15:04:05 MST\n"))
}

func main() {
	port := ":9999"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/time", timeHandler)

	http.ListenAndServe(port, nil)
}
