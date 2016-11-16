package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var cases = []struct {
	in, out string
}{
	{"1ambda@golang.org", "gopher 1ambda"},
	{"something", "dear something"},
}

func TestHandler(t *testing.T) {

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)

		if err != nil {
			t.Fatalf("could not create request %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected 200; got %d", rec.Code)
		}

		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response %q", rec.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			req, err := http.NewRequest(
				http.MethodGet,
				"http://localhost:8080/"+c.in,
				nil,
			)

			if err != nil {
				b.Fatalf("could not create request %v", err)
			}

			rec := httptest.NewRecorder()
			handler(rec, req)

			if rec.Code != http.StatusOK {
				b.Errorf("expected 200; got %d", rec.Code)
			}

			if !strings.Contains(rec.Body.String(), c.out) {
				b.Errorf("unexpected body in response %q", rec.Body.String())
			}
		}
	}
}
