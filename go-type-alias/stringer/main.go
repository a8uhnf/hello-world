package main

import "fmt"

type stringer interface {
	fmt.Stringer
}

func main() {
	var ms stringer
	var s fmt.Stringer
	fmt.Println(ms == s)
	ms = s
	s = ms
}
