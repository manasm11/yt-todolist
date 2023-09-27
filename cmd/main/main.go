package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveIndex)
	addr := ":8080"
	fmt.Println("Starting server at", addr)
	http.ListenAndServe(addr, nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}
