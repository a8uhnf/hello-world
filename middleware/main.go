package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne", r.Method)
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	finalHandler := http.HandlerFunc(final)

	http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	if err := http.ListenAndServe(":3031", nil); err != nil {
		panic(err)
	}
}
