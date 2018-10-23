package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/spf13/viper/remote"
)

// AllPods ...
type AllPods struct {
	Pods []string `json:"pods"`
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
}

func getPods() []string {
	return []string{"hello", "hello1", "world1"}
}

func main() {
	log.Println("--------------------")
	http.Handle("/", GetPodName{})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
