package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello World!!!")
	url := fmt.Sprintf("%s:%s/hello", "http://localhost", "8080") //:8080/hello
	_, err := http.Get(url)
	if err != nil {
		// panic(err)
		log.Println("error occured", err)
	}
	/* b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
	} */

	// logrus.Println("-------", string(b))
	for {

	}
}
