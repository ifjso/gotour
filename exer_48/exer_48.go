package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)

	for {
		pz := z
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))

		if cmplx.Abs(pz-z) < 0.000000001 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Cbrt(29))
}
