package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

//func Scale(v Vertex3, f float64) {
func Scale(v *Vertex3, f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func Abs2(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex3{3, 4}
	//Scale(v, 10)
	Scale(&v, 10)
	fmt.Println(Abs2(v))
}
