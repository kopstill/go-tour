package main

import (
	"fmt"
	"math"
)

const epison = 0.000001

func Sqrt(x float64) float64 {
	result := 1.
	for {
		result = result + (x/(2*result) - result/2)
		if math.Abs(result*result-x) < epison {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}
