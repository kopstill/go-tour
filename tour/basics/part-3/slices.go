package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = primes[1:4]
	//var s = primes[1:]
	//var s = primes[:4]
	//var s = primes[:]
	var s = primes[1:4]
	fmt.Println(s)
}
