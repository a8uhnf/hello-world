package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
	fmt.Println(viper.Get("port"))     // 8080
	fmt.Println(viper.Get("hostname")) // myhostname.com
}

func getPods() []string {
	nextHop := viper.Get("hostname").(string)
	if nextHop == "" {
		return []string{}
	}
	resp, err := http.Get(nextHop)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	pods := []string{}

	err = json.Unmarshal(b, &pods)
	if err != nil {
		panic(err)
	}
	return pods
	//return []string{"hello", "hello1", "world1"}
}

func main() {
	log.Println("--------------------")

	http.Handle("/foo", FooHandlerTest{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.Handle("/pod-name", GetPodName{})
	initConsul()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initConsul() {
	consulURL := os.Getenv("CONSUL_URL")
	if consulURL == "" {
		consulURL = "35.221.128.208:8500"
	}
	consulPATH := os.Getenv("CONSUL_PATH")
	if consulPATH == "" {
		consulPATH = "deploy-1"
	}

	viper.AddRemoteProvider("consul", consulURL, consulPATH)
	viper.SetConfigType("yaml") // Need to explicitly set this to json
	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("port"))     // 8080
	fmt.Println(viper.Get("hostname")) // myhostname.com
}
