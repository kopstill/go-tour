package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//func (v Vertex7) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

type Vertex7 struct {
	X, Y float64
}

func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex7{3, 4}

	a = f
	a = &v
	//a = v

	fmt.Println(a.Abs())
}
