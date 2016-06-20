package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] = %d\n", i, p[i])
	}

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}
