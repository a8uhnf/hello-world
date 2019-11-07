package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("hello world!!!")

	router := chi.NewRouter()

	router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		log.Println("---------------")
		q := r.URL.Query()
		if len(q["assigned"]) == 0 {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte("need to provide 'assined' query parameter "))
		}

		w.Write([]byte(fmt.Sprintf("%v", q["assigned"])))
	})

	if err := http.ListenAndServe(":8088", router); err != nil {
		panic(err)
	}
}
