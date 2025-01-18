package main

import (
	"log"
	"net/http"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	http.HandleFunc("/", serveIndex)
	//http.HandleFunc("/ws", nil)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
