package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
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
	err := viper.AddRemoteProvider("consul", consulURL, consulPATH)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}
