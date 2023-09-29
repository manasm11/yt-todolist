package server_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manasm11/yt-todolist/pkg/server"
	"github.com/manasm11/yt-todolist/pkg/todo"
)

var s = httptest.NewServer(server.NewTodoApiServeMux())

func TestServerEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, endpoint, method string
		statusCode             int
		checkResponseContent   func(io.ReadCloser) bool
	}{
		{"all todos", "/api/todo/", http.MethodGet, 200, checkTodosList},
		{"get one todo", "/api/todo/1/", http.MethodGet, 200, checkTodo},
		{"create todo", "/api/todo", http.MethodPost, 201, checkTodo},
		{"update todo", "/api/todo/1", http.MethodPut, 200, checkTodo},
		{"delete todo", "/api/todo/1/", http.MethodDelete, 200, nil},
		{"incorrect endpoint", "/wrongendpoint", http.MethodGet, 404, nil},
		{"incorrect enpoint starting with /api", "/api/wrongendpoint", http.MethodGet, 404, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, s.URL+tt.endpoint, nil)
			ok(t, err)
			res, err := http.DefaultClient.Do(req)
			ok(t, err)
			assertStatus(t, res, tt.statusCode)
			if tt.checkResponseContent != nil {
				assert(t, tt.checkResponseContent(res.Body), "incorrect response content: %v", res.Body)
			}
		})
	}

}

func checkTodosList(body io.ReadCloser) bool {
	return json.NewDecoder(body).Decode(&[]todo.Todo{}) == nil
}

func checkTodo(body io.ReadCloser) bool {
	return json.NewDecoder(body).Decode(&todo.Todo{}) == nil
}

func ok(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertStatus(t testing.TB, res *http.Response, want int) {
	t.Helper()
	assert(t, res.StatusCode == want, "expected status code %d, got %q", want, res.Status)
}

func assert(t testing.TB, condition bool, message string, args ...interface{}) {
	t.Helper()
	if condition != true {
		t.Fatalf(message, args...)
	}
}
