package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	r := http.RedirectHandler("https://google.com", 307)

	mux.Handle("/", r)

	http.ListenAndServe(":3030", mux)
}
