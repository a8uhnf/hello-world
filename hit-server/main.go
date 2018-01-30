package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Client Started....")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://192.168.1.104:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	for {
		fmt.Println("---")
	}
}
