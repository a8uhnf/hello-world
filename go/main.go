package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server Started....")
	http.HandleFunc("/hello", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
