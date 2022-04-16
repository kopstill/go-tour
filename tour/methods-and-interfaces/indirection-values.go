package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func (v Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex5{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex5{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
