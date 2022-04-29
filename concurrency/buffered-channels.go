package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	i := <-ch
	j := <-ch
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(j)
	fmt.Println(&j)
}
