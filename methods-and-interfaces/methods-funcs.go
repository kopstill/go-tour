package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

func Abs(vertex Vertex1) float64 {
	return math.Sqrt(vertex.X*vertex.X + vertex.Y*vertex.Y)
}

func main() {
	v := Vertex1{3, 4}
	fmt.Println(Abs(v))
}
