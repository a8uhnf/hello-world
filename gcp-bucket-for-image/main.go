package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const signerUrl = "http://hello.world.com:8081/sign"

func getSignedURL(target string, values url.Values) (string, error) {
	resp, err := http.PostForm(target, values)
	if err != nil {
		fmt.Println("xxx")
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func main() {
	// Get signed url from the API server hosted on App Engine.
	u, err := getSignedURL(signerUrl, url.Values{"content_type": {"image/png"}})
	if err != nil {
		fmt.Println("-------------")
		log.Fatal(err)
	}
	fmt.Printf("Signed URL here: %q\n", u)

	b, err := ioutil.ReadFile("./xyz.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// Generates *http.Request to request with PUT method to the Signed URL.
	req, err := http.NewRequest("PUT", u, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "image/png")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
