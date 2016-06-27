package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x

	for {
		nz := z - (z*z-x)/(2*z)

		if math.Abs(nz-z) < 0.00000001 {
			return nz
		}

		z = nz
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
