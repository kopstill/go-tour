package main

import (
	"fmt"
)

func fibonacci(cap int, c chan int) {
	x, y := 0, 1
	for i := 0; i < cap; i++ {
		c <- x
		x, y = y, x+y
		//time.Sleep(1000 * time.Millisecond)
		//fmt.Println("x")
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	//fibonacci(cap(c), c)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
