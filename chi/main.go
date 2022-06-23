package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var _ http.ResponseWriter = (*writer)(nil)

type writer struct {
	statusCode int
	http.ResponseWriter
}

// Header implements http.ResponseWriter
func (w *writer) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write implements http.ResponseWriter
func (w *writer) Write(in []byte) (int, error) {
	return w.ResponseWriter.Write(in)
}

// WriteHeader implements http.ResponseWriter
func (w *writer) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := &writer{ResponseWriter: w}
			next.ServeHTTP(ww, r)
			fmt.Printf("status code:%d\n", ww.statusCode)
		}
		return http.HandlerFunc(fn)
	})
	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)
			fmt.Printf("status code:%d\n", ww.Status())
		}
		return http.HandlerFunc(fn)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("welcome"))
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}
}
