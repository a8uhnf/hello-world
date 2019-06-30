package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func middlewareOne(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("----", r.Method)
		fn.ServeHTTP(w, r)
		log.Println("xxxxx", r.Method)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(middlewareOne)
	r.Method("GET", "/", Handler(customHandler))
	http.ListenAndServe(":3333", r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}
	log.Println("Hello world!!")
	w.Write([]byte("foo"))
	return nil
}
