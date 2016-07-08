package main

import "fmt"

func fibonacci() func() int {
	n1, n2 := 0, 1
	return func() (n int) {
		n, n1, n2 = n1, n2, n1+n2
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
