package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		pz := z
		z = z - (z*z-x)/(2*z)

		if math.Abs(pz-z) < 0.00000001 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
