package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sirupsen/logrus"
)

func main() {
	fmt.Println("hello World!!!")
	url := fmt.Sprintf("%s:%s/hello", "http://172.17.0.2", "8080") //:8080/hello
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error occured", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error occured....", err)
	}

	logrus.Println("-------", string(b))
	for {

	}
}
