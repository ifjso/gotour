package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os)
	}
}
