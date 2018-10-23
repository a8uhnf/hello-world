package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

// AllPods ...
type AllPods struct {
	Pods []string `json:"pods"`
}

// FooHandlerTest ...
type FooHandlerTest struct{}

// ServerHTTP ...
func (foo FooHandlerTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// GetPodName ...
type GetPodName struct{}

func (pod GetPodName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pods := getPods()
	podJSON, err := json.Marshal(pods)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(podJSON)
	// fmt.Fprintf(w, "hello, %q", os.Getenv("POD_NAME"))
}

func getPods() []string {
	return []string{"hello", "hello1", "world1"}
}

func main() {
	fmt.Println("--------------------")
	log.Println("--------------------")

	http.Handle("/foo", FooHandlerTest{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.Handle("/pod-name", GetPodName{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
