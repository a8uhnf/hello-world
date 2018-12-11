package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("----------------------")
	consulURL := os.Getenv("CONSUL_URL")
	if consulURL == "" {
		consulURL = "localhost:8502"
	}
	consulPATH := os.Getenv("CONSUL_PATH")
	if consulPATH == "" {
		consulPATH = "test-helm"
	}
}
