package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	log.Println(r.Method)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(errors.Wrap(fmt.Errorf("test"), "Hello World!!! Testing Error!!!"))
	log.Println(string(b))
	log.Println(w.Header())
}

func main() {
	fmt.Println("Server Started....")
	http.HandleFunc("/hello", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
