package main

import (
	"fmt"
	"math"
)

const EPSILON = 0.00000001

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		nz := z - (z*z-x)/(2*z)

		if math.Abs(nz-z) < EPSILON {
			return nz
		}

		z = nz
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
