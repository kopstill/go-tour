package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

//func (v Vertex2) Scale(f float64) {
func (v *Vertex2) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func (v Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs2())
}
