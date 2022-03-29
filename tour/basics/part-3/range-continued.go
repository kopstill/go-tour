package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = i << uint(i)
	}
	fmt.Println(pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
