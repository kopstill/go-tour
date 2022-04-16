package main

import "fmt"

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func ScaleFunc(v *Vertex4, f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func main() {
	v := Vertex4{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex4{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
