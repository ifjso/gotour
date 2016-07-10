package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var fnWalk func(*tree.Tree, chan int)
	fnWalk = func(t *tree.Tree, ch chan int) {
		if t != nil {
			fnWalk(t.Left, ch)
			ch <- t.Value
			fnWalk(t.Right, ch)
		}
	}

	fnWalk(t, ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
