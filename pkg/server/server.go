package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/manasm11/yt-todolist/pkg/todo"
)

type todoApiServeMux struct {
	r  http.Handler
	td todo.TodoDao
}

func NewTodoApiServeMux() (s *todoApiServeMux) {
	s = new(todoApiServeMux)
	s.r = http.NewServeMux()
	// s.td = tododao
	return
}

func (s todoApiServeMux) handleAllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]todo.Todo{})
}

func (s todoApiServeMux) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo.Todo{})
}

func (s todoApiServeMux) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todo.Todo{})
}

func (s todoApiServeMux) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todo.Todo{})
}

func (s todoApiServeMux) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]todo.Todo{})
}

func (s todoApiServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := withoutTrailingSlash(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		if path == "/api/todo" {
			s.handleAllTodos(w, r)
			return
		}
		if strings.HasPrefix(path, "/api/todo/") {
			s.handleGetTodo(w, r)
			return
		}
	case http.MethodPost:
		if path == "/api/todo" {
			s.handleCreateTodo(w, r)
			return
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/api/todo/") {
			s.handleUpdateTodo(w, r)
			return
		}
	case http.MethodDelete:
		if strings.HasPrefix(path, "/api/todo/") {
			s.handleDeleteTodo(w, r)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}

func withoutTrailingSlash(rawPath string) string {
	if strings.HasSuffix(rawPath, "/") {
		return rawPath[:len(rawPath)-1]
	}
	return rawPath
}
