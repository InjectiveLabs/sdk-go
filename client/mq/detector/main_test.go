package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestControlPlaneBlocksAddsBearerToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}

		if r.URL.Path != "/request" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		if got := r.URL.Query().Get("from_height"); got != "42" {
			t.Fatalf("unexpected from_height: %q", got)
		}

		if got := r.Header.Get("Authorization"); got != "Bearer secret" {
			t.Fatalf("unexpected Authorization header: %q", got)
		}

		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	err := requestControlPlaneBlocks(context.Background(), server.Client(), server.URL, 42, "secret")
	if err != nil {
		t.Fatalf("requestControlPlaneBlocks returned error: %v", err)
	}
}

func TestRequestControlPlaneBlocksOmitsBearerTokenWhenEmpty(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "" {
			t.Fatalf("unexpected Authorization header: %q", got)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := requestControlPlaneBlocks(context.Background(), server.Client(), server.URL, 42, "")
	if err != nil {
		t.Fatalf("requestControlPlaneBlocks returned error: %v", err)
	}
}
