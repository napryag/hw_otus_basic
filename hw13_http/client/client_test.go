package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequest(t *testing.T) {
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer ts.Close()

	resp, err := sendRequest(ctx, "GET", ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestBuildURL(t *testing.T) {
	tests := []struct {
		serverURL string
		path      string
		expected  string
	}{
		{"http://localhost:8080", "/test", "http://localhost:8080/test"},
		{"localhost:8080", "/test", "http://localhost:8080/test"},
	}

	for _, tc := range tests {
		result, err := buildURL(tc.serverURL, tc.path)
		if err != nil {
			t.Fatalf("buildURL failed: %v", err)
		}
		if result != tc.expected {
			t.Errorf("buildURL(%q, %q) = %q; want %q", tc.serverURL, tc.path, result, tc.expected)
		}
	}
}
