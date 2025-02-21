package query2port_test

import (
	"context"
	"github.com/brudnevskij/query2port"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryPortForwarder(t *testing.T) {
	cfg := query2port.CreateConfig()
	cfg.QueryParamName = "port"

	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handler, err := query2port.New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:223/test?port=1337", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)
	if req.URL.Host != "localhost:1337" || req.URL.Port() != "1337" {
		t.Fatalf("expected: localhost:1337, got: %s", req.URL.Host)
	}

	req = httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/test?port=7777", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)
	if req.URL.Host != "localhost:7777" || req.URL.Port() != "7777" {
		t.Fatalf("expected: localhost:7777, got: %s", req.URL.Host)
	}

	req = httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:2234/test?port=", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)
	if req.URL.Host != "localhost:2234" || req.URL.Port() != "2234" {
		t.Fatalf("expected: localhost:2234, got: %s", req.URL.Host)
	}

	req = httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1234/test", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req)
	if req.URL.Host != "localhost:1234" || req.URL.Port() != "1234" {
		t.Fatalf("expected: localhost:7777, got: %s", req.URL.Host)
	}
}
