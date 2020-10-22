package main

import (
	"fmt"

	"github.com/uber/h3-go"
)

func main() {
	ExampleFromGeo()
}
func ExampleFromGeo() {
	geo := h3.GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9
	fmt.Println("-----------------------")
	fmt.Printf("%#x\n", h3.FromGeo(geo, resolution))
	// Output:
	// 0x8928308280fffff
}
