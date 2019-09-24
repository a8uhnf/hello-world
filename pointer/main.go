package main

import (
	"fmt"
	"sync"
	"time"
)

type test interface {
	Hello() string
	World() string
}

type earth struct {
	A int
	B int
}

func (e *earth) Hello() string {
	return "hello"
}

func (e *earth) World() string {
	return "earth"
}

func main() {
	fmt.Println("-----")
	mp := make(map[string]test)
	mp["1"] = new(earth)

	wg := sync.WaitGroup{}
	it := 3
	wg.Add(it)

	for i := 0; i < it; i++ {
		go func(yo int) {
			wg.Wait()
			t := time.Now().String()
			e := mp["1"].(*earth)
			e.A = yo
			e.B = yo
			fmt.Printf("time: %v: %v %v\n", t, e.A, e.B)

		}(i)
		wg.Done()
	}
	ch := make(chan struct{})
	<-ch
}
