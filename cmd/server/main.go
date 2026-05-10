package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/ryan-ching/kv-store/internal/store"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP server address")
	flag.Parse()

	s := store.New()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /kv/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		value, ok := s.Get(key)
		// TODO: if !ok, return 404. Otherwise write `value` to w.
		//       Decide: what Content-Type makes sense if values are arbitrary bytes?
		_, _ = value, ok
		http.Error(w, "not implemented", http.StatusNotImplemented)
	})

	mux.HandleFunc("PUT /kv/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// TODO: store `body` under `key` and return 204.
		//       Should you cap the body size? See http.MaxBytesReader.
		s.Put(key, body)
		http.Error(w, "not implemented", http.StatusNotImplemented)
	})

	mux.HandleFunc("DELETE /kv/{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		// TODO: return 204 — or 404 if the key wasn't there? Pick one, justify it in questions.md.
		s.Delete(key)
		http.Error(w, "not implemented", http.StatusNotImplemented)
	})

	log.Printf("listening on %s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
