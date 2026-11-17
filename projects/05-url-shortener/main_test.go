package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShortenAndResolve(t *testing.T) {
	store := NewStore()
	handler := newHandler(store)

	body := `{"url":"https://example.com/page"}`
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(body))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d; body = %s", rec.Code, rec.Body.String())
	}
	var resp shortenResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Code == "" {
		t.Fatal("expected a code")
	}

	req = httptest.NewRequest(http.MethodGet, "/"+resp.Code, nil)
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusFound {
		t.Fatalf("redirect status = %d", rec.Code)
	}
	if got := rec.Header().Get("Location"); got != "https://example.com/page" {
		t.Errorf("Location = %q", got)
	}
}

func TestNotFound(t *testing.T) {
	handler := newHandler(NewStore())
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d; want 404", rec.Code)
	}
}

func TestBadURL(t *testing.T) {
	handler := newHandler(NewStore())
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(`{"url":"not a url"}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d; want 400", rec.Code)
	}
}

func TestStoreConcurrent(t *testing.T) {
	store := NewStore()
	done := make(chan struct{})
	for i := 0; i < 50; i++ {
		go func() {
			code := store.Shorten("https://example.com")
			store.Resolve(code)
			done <- struct{}{}
		}()
	}
	for i := 0; i < 50; i++ {
		<-done
	}
	if store.Count() != 50 {
		t.Errorf("count = %d; want 50", store.Count())
	}
}
