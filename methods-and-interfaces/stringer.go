package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old).", p.Name, p.Age)
}

func main() {
	a := Person{"Ada", 42}
	t := Person{"Tom", 9001}
	fmt.Println(a, t)
}
