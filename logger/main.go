package main

import (
	"github.com/Sirupsen/logrus"
	"log"
	"time"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		log.Println("Hello World")
		log.Println("hello")
		logrus.Println("Hello Hanifa....")
		log.Println("time", time.Now())
	}
}
