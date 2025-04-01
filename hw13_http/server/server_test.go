package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		method       string
		path         string
		expectedBody string
	}{
		{"GET", "/test", "Server received GET request at /test"},
		{"POST", "/test", "Server received POST request at /test"},
	}

	for _, tc := range tests {
		req, err := http.NewRequestWithContext(ctx, tc.method, tc.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != tc.expectedBody {
			t.Errorf("handler returned unexpected body: got %q want %q", rr.Body.String(), tc.expectedBody)
		}
	}
}
