package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

func getSignedurl() {
	sakey, err := ioutil.ReadFile("../xyz.json")
	if err != nil {
		// TODO: handle error.
		panic(err)
	}
	fmt.Printf("%s", sakey)
	cfg, err := google.JWTConfigFromJSON(sakey)
	if err != nil {
		panic(err)
	}
	url, err := storage.SignedURL("xyz-bucket", "xyz.jpg", &storage.SignedURLOptions{
		GoogleAccessID: cfg.Email,
		PrivateKey:     cfg.PrivateKey,
		Method:         "PUT",
		Expires:        time.Now().Add(48 * time.Hour),
		ContentType:    "image/jpeg",
	})
	if err != nil {
		// TODO: handle error.
		panic(err)
	}
	fmt.Println(url)
}

func main() {
	getSignedurl()
}
