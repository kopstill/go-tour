package main

import (
	"fmt"
	"strconv"
)

func fibonacci1(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			//default:
			//	fmt.Println("default")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		j := <-c
		fmt.Println(strconv.Itoa(j) + "---")
		quit <- 0
		//fmt.Println(quit)
	}()
	fibonacci1(c, quit)
}
