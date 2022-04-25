package main

import "fmt"

type I3 interface {
	M()
}

func main() {
	var i I3
	describe2(i)
	i.M()
}

func describe2(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
