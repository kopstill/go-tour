package main

import (
	"fmt"
	"math"
)

func main() {
	test()

	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)
}

func test() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	i1 := 42
	f1 := float64(i1)
	u1 := uint(f1)
	fmt.Println(i1, f1, u1)
}
