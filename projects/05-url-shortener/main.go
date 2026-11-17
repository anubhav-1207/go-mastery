package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func newHandler(store *Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/shorten", func(w http.ResponseWriter, r *http.Request) {
		var req shortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		parsed, err := url.ParseRequestURI(req.URL)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
			http.Error(w, "invalid URL", http.StatusBadRequest)
			return
		}
		code := store.Shorten(req.URL)
		writeJSON(w, http.StatusCreated, shortenResponse{Code: code, URL: req.URL})
	})

	mux.HandleFunc("GET /{code}", func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		target, err := store.Resolve(code)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.Error(w, "not found", http.StatusNotFound)
			} else {
				http.Error(w, "error", http.StatusInternalServerError)
			}
			return
		}
		http.Redirect(w, r, target, http.StatusFound)
	})

	mux.HandleFunc("GET /api/stats", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]int{"count": store.Count()})
	})

	return requestID(mux)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("encode:", err)
	}
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = encode(int(time.Now().UnixNano()))
		}
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r)
	})
}

func main() {
	store := NewStore()
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           newHandler(store),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("url shortener listening on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
