package main

import (
	"fmt"
	"net/http"
)

type Employee interface {
	Read() string
	HelloWorld() string
}

type Hanif struct{}

func (h Hanif) Read() string {
	return "hello from hanif"
}
func (h Hanif) HelloWorld() string {
	return "hello from hanif"
}

type Binoy struct{}

func (h Binoy) Read() string {
	return "hello from binoy"
}

func (h Binoy) HelloWorld() string {
	return "hello from binoy"
}

var mp map[string]Employee

func main() {
	var a interface{}
	mp = make(map[string]Employee)
	mp["/hanifa"] = new(Hanif)
	mp["/binoy"] = new(Binoy)
	a = "1"
	// /sticker-diver/communicate/{employee}

	http.HandleFunc("/", handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	// w.Write([]byte("hello hanifa"))
	if r.URL.String() == "/hanifa" {
		v := mp[r.URL.String()]
		w.Write([]byte(v.Read()))
	} else if r.URL.String() == "/binoy" {
		v := mp[r.URL.String()]
		w.Write([]byte(v.Read()))
	}
}
