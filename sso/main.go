package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Hnaifa. you dumb ass......", os.Getenv("APP"))
	fmt.Fprintf(w, "Hi there, I love %s! %s", r.URL.Path[1:], os.Getenv("APP"))
}

func main() {
	fmt.Println("Server Started....")
	http.HandleFunc("/hello", handler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
