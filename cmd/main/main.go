package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/manasm11/yt-todolist/pkg/server"
)

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		server.NewTodoApiServeMux().ServeHTTP(w, r)
	})
	http.HandleFunc("/", serveIndex)
	addr := ":8080"
	fmt.Println("Starting server at", addr)
	http.ListenAndServe(addr, nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	r.URL.RawPath = strings.TrimPrefix(r.URL.RawPath, "/api")
	http.ServeFile(w, r, "html/index.html")
}
