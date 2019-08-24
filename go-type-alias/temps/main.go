package main

import "fmt"

type hello float64

func (h hello) String() string {
	return fmt.Sprintf("%.2f C", h)
}

func main() {
	h := hello(10)
	fmt.Println(h)
}
