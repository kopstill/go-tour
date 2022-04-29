package main

import (
	"fmt"
	"strconv"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s + strconv.Itoa(i))
	}
}

func main() {
	go say("world")
	say("hello")
}
