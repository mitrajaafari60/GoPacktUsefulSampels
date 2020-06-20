package main

import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " d1 ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " d2 ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " d3 ")
		}(i)
	}
}
func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
