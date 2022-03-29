package main

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var map1 = map[string]Vertex5{
	"Bell Labs": Vertex5{
		40.68433, -74.39967,
	},
	"Google": Vertex5{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(map1)
}
