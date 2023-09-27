package server

import (
	"net/http"

	"github.com/manasm11/yt-todolist/pkg/todo"
)

type TodoApiServeMux struct {
	td todo.TodoDao
}

func (s TodoApiServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
