package main

import (
	"fmt"
	"strconv"
)

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	i, j := 1, 2
	i, j = 3, 4
	i = 5
	j = 6
	//var s = "hello"
	//s := "hello"
	//var s = string(123)
	s := strconv.Itoa(123)
	s1 := fmt.Sprint(123.456)
	fmt.Println(i, j, s, s1)

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
