package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't sqrt negative number: %v", float64(e))
	//return fmt.Sprintf("can't sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func main() {
	var v float64 = -2
	r, e := Sqrt(v)
	if e == nil {
		fmt.Printf("sqrt number %v is %v", v, r)
	} else {
		fmt.Println(e)
	}
}
