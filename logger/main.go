package main

import (
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		log.Println("Hello World")
		log.Println("hello")
		log.Println("time", time.Now())
	}
}
