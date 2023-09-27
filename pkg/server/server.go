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
	s.r.HandlerFunc(http.MethodGet, "/todos/", s.handleAllTodos)
	s.r.HandlerFunc(http.MethodPost, "/todo/", s.handleCreateTodo)
	s.r.HandlerFunc(http.MethodGet, "/todo/:id", s.handleGetTodo)
	s.r.HandlerFunc(http.MethodPut, "/todo/:id", s.handleUpdateTodo)
	s.r.HandlerFunc(http.MethodDelete, "/todo/:id", s.handleDeleteTodo)
	// s.td = tododao
	return
}

func (s todoApiServeMux) handleAllTodos(w http.ResponseWriter, r *http.Request) {

}

func (s todoApiServeMux) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (s todoApiServeMux) handleGetTodo(w http.ResponseWriter, r *http.Request) {

}

func (s todoApiServeMux) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func (s todoApiServeMux) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func (s todoApiServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
