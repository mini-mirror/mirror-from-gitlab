package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestMock(t *testing.T) {
	expected := "Hello!"
	actual := "Hello!"
	if actual != expected {
		t.Errorf("Test failed", expected, actual)
	}
}

func TestRunServe(t *testing.T) {
	router := httprouter.New()

	routed := false

	router.Handle("GET", "/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		routed = true
		want := httprouter.Params{httprouter.Param{"name", "Eleanore"}}
		if !reflect.DeepEqual(ps, want) {
			t.Fatalf("wrong wildcard values: want %v, got %v", want, ps)
		}
	})

	w := new(mockResponseWriter)

	req, _ := http.NewRequest("GET", "/hello/Eleanore", nil)
	router.ServeHTTP(w, req)

	if !routed {
		t.Fatal("routing failed!")
	}

}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}
