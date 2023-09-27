package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manasm11/yt-todolist/pkg/server"
)

func TestServerEndpoints(t *testing.T) {
	tests := []struct {
		name, endpoint, method string
		statusCode             int
	}{
		{"all todos", "/todo/", http.MethodGet, 200},
	}

	s := httptest.NewServer(server.TodoApiServeMux{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, s.URL+tt.endpoint, nil)
			ok(t, err)
			res, err := http.DefaultClient.Do(req)
			ok(t, err)
			assertStatus(t, res, tt.statusCode)
		})
	}
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