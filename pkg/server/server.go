package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manasm11/yt-todolist/pkg/todo"
)

type todoApiServeMux struct {
	r  *httprouter.Router
	td todo.TodoDao
}

func NewTodoApiServeMux() (s *todoApiServeMux) {
	s = new(todoApiServeMux)
	s.r = httprouter.New()
	s.r.HandlerFunc(http.MethodGet, "/todos/", func(w http.ResponseWriter, r *http.Request) {})
	s.r.HandlerFunc(http.MethodPost, "/todo/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) })
	s.r.HandlerFunc(http.MethodGet, "/todo/:id", func(w http.ResponseWriter, r *http.Request) {})
	s.r.HandlerFunc(http.MethodPut, "/todo/:id", func(w http.ResponseWriter, r *http.Request) {})
	s.r.HandlerFunc(http.MethodDelete, "/todo/:id", func(w http.ResponseWriter, r *http.Request) {})
	// s.td = tododao
	return
}

func (s todoApiServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
